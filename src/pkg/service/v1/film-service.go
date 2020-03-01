package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/tuanpa28/film-service/src/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)


type FilmServiceServer struct {
	database *sql.DB
	savedCategories [] *v1.FilmCategory

	mu sync.Mutex
}


func (s *FilmServiceServer) ListFashionCategories(empty *v1.FilmEmptyRequest,stream v1.FilmsService_ListFashionCategoriesServer) error {
	for _, category := range  s.savedCategories {
		if err := stream.Send(category); err != nil {
			return err
		}
	}

	return nil
}

func (s * FilmServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	conn, err := s.database.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to connect database -> Error :" + err.Error())
	}

	return conn, nil
}

func NewFilmServiceServer(database *sql.DB) v1.FilmsServiceServer {
	var service = FilmServiceServer{database: database}
	service.loadLocalDataForSavedCategories()
	return &service
}

func (s *FilmServiceServer) SendMessage(stream v1.FilmsService_SendMessageServer) error {
	for {
		in, err := stream.Recv()
		// end of the streaming
		if err == io.EOF {
			grpclog.Println("server -- finished stream")
			return nil
		}
		if err != nil {
			grpclog.Printf("returned with error %v", err)
			return err
		}
		message := in.Message
		name := in.Name
		if name == "" {
			name = "Unknown"
		}
		grpclog.Printf("server -- received message:\n%v: %v", name, message)
		revMsg := "received" + time.Now().Format("2006-01-02 15:04:05")
		stream.Send(&v1.ChatStreamResponse{
			Message: revMsg,
		})
	}
}

func (s * FilmServiceServer) loadLocalDataForSavedCategories() {

	jsonFile, err := os.Open("./src/pkg/service/v1/categories.db.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &s.savedCategories)
	defer jsonFile.Close()
}


