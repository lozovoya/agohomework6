package main

import (
	"agohomework6/cmd/server/app"
	bankgrpcv1 "agohomework6/pkg/bank/v1"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const defaultPort = "9999"
const defaultHost = "0.0.0.0"
const defaultDb = "postgres://app:pass@localhost:5432/db"

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	db, ok := os.LookupEnv("DB")
	if !ok {
		db = defaultDb
	}

	if err := execute(net.JoinHostPort(host, port), db); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(addr string, db string) error {

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, db)
	if err != nil {
		return err
	}
	defer pool.Close()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	server := app.NewServer(pool, ctx)
	bankgrpcv1.RegisterTemplateServiceServer(grpcServer, server)

	return grpcServer.Serve(listener)
}
