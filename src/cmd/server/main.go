package main

import (
	"fmt"
	cmd "github.com/tuanpa28/film-service/src/pkg/cmd"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		GRPCPort string `yaml:"grpcPort"`
		HttpPort string `yaml:"httpPort"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Host string  `yaml:"host"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Schema string `yaml:"schema"`
	} `yaml:"database"`
}

func main() {

	file, err := os.Open("./src/configs/config.yml")

	if err != nil {
		fmt.Println("Can not loaded configs from file")
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Can not decoder data from file")
	}



	if err :=  cmd.RunServer(cmd.Config{
		GRPCPort:            cfg.Server.GRPCPort,
		HTTPPort:            cfg.Server.HttpPort,
		DataStoreDBHost:     cfg.Database.Host,
		DataStoreDBUser:     cfg.Database.Username,
		DataStoreDBPassword: cfg.Database.Password,
		DataStoreDBSchema:   cfg.Database.Schema,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
