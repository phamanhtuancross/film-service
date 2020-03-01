package v1

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tuanpa28/film-service/src/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"sync"
)


type FilmServiceServer struct {
	database *sql.DB
	jwtPrivatekey *rsa.PrivateKey
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

	key, _ := ioutil.ReadFile("./src/configs/jwt/private.pem")
	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		print(err)
	}
	service.jwtPrivatekey = parsedKey
	return &service
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
