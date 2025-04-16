package main

import (
	"context"
	"log"
	"net"
	"s0611-23/internal/demobase"
	"s0611-23/internal/proxyproto"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

const connStr = "postgres://postgresuser:postgrespass@127.0.0.1:5432/demobase?sslmode=disable"

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalln(err)
	}

	querier := demobase.New(conn)

	svc := NewService(querier)

	proxyproto.RegisterCentrifugoProxyServer(srv, svc)

	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
