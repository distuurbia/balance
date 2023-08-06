// Package main contains main goroutine and connect to postgress function
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/caarlos0/env"
	"github.com/distuurbia/balance/internal/config"
	"github.com/distuurbia/balance/internal/handler"
	"github.com/distuurbia/balance/internal/repository"
	"github.com/distuurbia/balance/internal/service"
	protocol "github.com/distuurbia/balance/protocol/balance"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func connectPostgres(cfg *config.Config) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(cfg.PostgresPath)
	if err != nil {
		return nil, fmt.Errorf("error in method pgxpool.ParseConfig: %v", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, fmt.Errorf("error in method pgxpool.NewWithConfig: %v", err)
	}
	return pool, nil
}

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("main -> %v", err)
	}
	pool, err := connectPostgres(&cfg)
	if err != nil {
		logrus.Fatalf("main -> %v", err)
	}
	validate := validator.New()
	r := repository.NewBalanceRepository(pool)
	s := service.NewBalanceService(r)
	h := handler.NewBalanceHandler(s, validate)
	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		logrus.Fatalf("main -> %v", err)
	}
	serverRegistrar := grpc.NewServer()
	protocol.RegisterBalanceServiceServer(serverRegistrar, h)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		logrus.Fatalf("main -> %v", err)
	}
}
