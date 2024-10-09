package main

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type httpServer struct {
	addr string
}

func NewHttpServer (addr string) *httpServer{
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error{

	router := http.NewServeMux()

	router.HandleFunc("POST /v1/optimize", handleOptimize)

	// Setup connection with grpc
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	log.Printf("Serving on port: %s\n", s.addr)
	return http.ListenAndServe(s.addr, router)
}