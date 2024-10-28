package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/fanfaronDo/anomaly_detection/internal/server"
	api "github.com/fanfaronDo/anomaly_detection/pkg/api/api/proto"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Filure creating server listener: %s", err)
	}

	serv := grpc.NewServer()

	api.RegisterDataServiceServer(serv, &server.Server{})
	log.Println("Server is running on port 8080...")
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
