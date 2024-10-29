package main

import (
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/fanfaronDo/anomaly_detection/pkg/api/api/proto"
	"github.com/fanfaronDo/anomaly_detection/pkg/config"
	"github.com/fanfaronDo/anomaly_detection/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	host := cfg.ServerTransmitterHost + ":" + cfg.ServerTransmitterPort
	rand.Seed(time.Now().UnixNano())

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Filure creating server listener: %s", err)
	}

	serv := grpc.NewServer()

	api.RegisterDataServiceServer(serv, &server.Server{})

	log.Printf("Server is running on %s...\n", host)

	defer lis.Close()

	go func() {
		if err := serv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Shutting down server...\n")
}
