package grpc

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	v1 "github.com/tuanpa28/film-service/src/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func RunServer(ctx context.Context, v1Api v1.FilmsServiceServer, port string) error {
	listener, err := net.Listen("tcp", ":" + port)


	if err != nil {
		return err
	}

	server := grpc.NewServer(grpc.StreamInterceptor(streamInterceptor))
	v1.RegisterFilmsServiceServer(server,  v1Api)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)


	go func() {
		for range c {
			log.Print("shutting down grpc server.........")
			server.GracefulStop()
			<- ctx.Done()
		}
	}()

	log.Println("starting grpc server..... on port"  + port)

	return server.Serve(listener)
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func validateToken(token string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			log.Printf("Unexpected signing method: %v", t.Header["alg"])
			return nil, fmt.Errorf("invalid token")
		}
		return publicKey, nil
	})
	if err == nil && jwtToken.Valid {
		return jwtToken, nil
	}
	return nil, err
}


func getRSAPublicKey() (*rsa.PublicKey, error) {
	key, _ := ioutil.ReadFile("./src/configs/jwt/public.pem")
	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	return parsedKey, err
}

// wrappedStream  wraps around the embedded grpc.ClientStream, and intercepts the RecvMsg and
// SendMsg method call.
type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	println("Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	println("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errMissingMetadata
	}

	jwtToken, ok := md["authorization"]
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "valid token required.")
	}

	rsaPublicKey, err := getRSAPublicKey()
	if err != nil {
		return err
	}
	_, err = validateToken(jwtToken[0], rsaPublicKey)
	if !ok {
		return errInvalidToken
	}

	err = handler(srv, newWrappedStream(ss))
	if err != nil {
		println("RPC failed with error %v", err)
	}
	return err

}
