package main

import (
	"fmt"
	"log"
	api "main/pkg/api/api/proto"
	"math/rand"
	"net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedDataServiceServer
}

func (serv *server) GenerateData(stream api.DataService_GenerateDataServer) error {
	mean := rand.Float64()*20 - 10
	stdDev := rand.Float64()*1.2 + 0.3
	sessionID := fmt.Sprintf("%s", uuid.New())

	for {
		frequency := rand.NormFloat64()*stdDev + mean
		timestamp := time.Now().Unix()

		entry := &api.DataEntry{
			SessionId: sessionID,
			Frequency: frequency,
			Timestamp: timestamp,
		}

		if err := stream.Send(entry); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Filure creating server listener: %s", err)
	}

	serv := grpc.NewServer()

	api.RegisterDataServiceServer(serv, &server{})
	log.Println("Server is running on port 8080...")
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
