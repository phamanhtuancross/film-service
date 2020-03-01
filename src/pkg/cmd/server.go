package cmd

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	grpc "github.com/tuanpa28/film-service/src/pkg/protocol/grpc"
	v1 "github.com/tuanpa28/film-service/src/pkg/service/v1"
)

type Config struct {
	GRPCPort            string
	HTTPPort			string
	DataStoreDBHost     string
	DataStoreDBUser     string
	DataStoreDBPassword string
	DataStoreDBSchema   string
}

func RunServer(config Config) error {

	ctx := context.Background()

	if len(config.GRPCPort) == 0 {
		return fmt.Errorf("Invalid gRPC Port for gRPC server '%s'", config.GRPCPort)
	}

	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		config.DataStoreDBUser,
		config.DataStoreDBPassword,
		config.DataStoreDBHost,
		config.DataStoreDBSchema,
		param)

	db,err := sql.Open("mysql", dsn)
	defer db.Close()

	if err != nil {
		fmt.Errorf("Failed to open database -> '%v'" , err)
	}

	var v1Api = v1.NewFilmServiceServer(db)
	return grpc.RunServer(ctx, v1Api, config.GRPCPort)
}


