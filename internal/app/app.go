package app

import (
	"flag"
	"fmt"
	"github.com/slavikx4/short-link/internal/contracts"
	"github.com/slavikx4/short-link/internal/database/inMemory"
	"github.com/slavikx4/short-link/internal/database/postgres"
	"github.com/slavikx4/short-link/internal/services"
	grpcShortLink "github.com/slavikx4/short-link/internal/transport/grpc"
	"github.com/slavikx4/short-link/pkg/api/proto"
	"github.com/slavikx4/short-link/pkg/logger"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

const (
	databaseURLKey = "DATABASE_URL"
	portKey        = "PORT"
)

func Run() {
	logger.Logger.Process.Println("запуск сервера")

	if err := initConfig(); err != nil {
		logger.Logger.Error.Fatalln("ошибка инициализации конфига: ", err)
	}

	//выбор БД через флаг
	//если true, значит postgres
	//по умолчанию false, значит inMemory
	var storage contracts.Storage
	if chooseStorage() {
		dbConfig := os.Getenv(databaseURLKey)
		if dbConfig == "" {
			logger.Logger.Error.Println("пустой env конфиг")
			dbConfig = viper.GetString(databaseURLKey)
		}

		pool, err := postgres.NewPoolPostgres(dbConfig)
		if err != nil {
			logger.Logger.Error.Fatalln("ошибка инициализации БД: ", err)
		}
		storage = postgres.NewStoragePostgres(pool)
	} else {
		storage = inMemory.NewStorageInMemory()
	}

	//создание сервиса бизнес слоя
	service := services.NewService(storage)

	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterShortLinkServer(server, &grpcShortLink.ServerShortLinkGRPC{Service: service})
	logger.Logger.Process.Println("создан сервер gRPC")

	logger.Logger.Process.Println("запуск листенера")
	port := viper.GetString(portKey)
	listener, err := net.Listen("tcp", fmt.Sprintf(":"+port))
	if err != nil {
		logger.Logger.Error.Fatalln("не получилось запустить листенер: ", err)
	}
	logger.Logger.Process.Printf("сервер прослушивается на порту: ", listener.Addr())

	if err = server.Serve(listener); err != nil {
		logger.Logger.Error.Fatalln("ошибка обслуживания сервера gRPC: ", err)
	}
}

func chooseStorage() bool {
	var dbFlag bool
	flag.BoolVar(&dbFlag, "storage", false, "run with storage Postgres(true/false)=")
	flag.Parse()
	return dbFlag
}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
