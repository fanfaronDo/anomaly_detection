package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fanfaronDo/anomaly_detection/internal/client"
	pb "github.com/fanfaronDo/anomaly_detection/pkg/api/api/proto"
	"github.com/fanfaronDo/anomaly_detection/pkg/config"
	repo "github.com/fanfaronDo/anomaly_detection/pkg/repository"

	"google.golang.org/grpc"
)

var (
	logfileName = "anomaly_detection.log"
	pool        = sync.Pool{
		New: func() interface{} {
			return &pb.DataEntry{}
		},
	}
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s -k <anomaly_coefficient>", os.Args[0])
	}

	flagK := flag.Float64("k", 0.0, "Anomaly coefficient")
	flag.Parse()

	myConfig, _ := config.LoadConfig()

	logFile, err := os.OpenFile(logfileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal("Error opening logfile %s", logfileName)
	}

	defer logFile.Close()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Receiver(*flagK, *myConfig, logFile, wg)
	wg.Wait()

}

func Receiver(k float64, cfg config.Config, logFile *os.File, wg *sync.WaitGroup) {
	statistics := &client.Statistics{}
	host := cfg.ServerTransmitterHost + ":" + string(cfg.ServerTransmitterPort)
	connectionToClient, err := grpc.Dial(host,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	defer wg.Done()

	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
		fmt.Errorf(logfileName, "Failed to connect to the server: %v\n", err)
		os.Exit(1)
	}
	defer connectionToClient.Close()
	log.Printf("Client connect to %s\n", host)

	defer func() {
		log.Printf("Reciver is closed\n")
		log.Printf("Client is closed\n")
	}()

	conn, err := repo.NewConnector(cfg)
	if err != nil {
		log.Fatalf("Errer connect to db %s\n", err)
		fmt.Fprintf(logFile, "Errer connect to db %s\n", err)
		os.Exit(1)
	}

	r := repo.NewRepository(conn)

	client := pb.NewDataServiceClient(connectionToClient)
	stream, err := client.GenerateData(context.Background())

	log.Printf("Reciver is running...\n")
	for {
		entry := pool.Get().(*pb.DataEntry)
		entry, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving data: %v", err)
			fmt.Fprintf(logFile, "error receiving data: %v\n", err)
		}

		if statistics.DetectAnomaly(entry.Frequency, k) {
			r.Create(*entry)
			log.Printf("Received Data: Session ID: %s, Frequency: %f, Timestamp: %d\n", entry.SessionId, entry.Frequency, entry.Timestamp)
			fmt.Fprintf(logFile, "Received Data: Session ID: %s, Frequency: %f, Timestamp: %d\n", entry.SessionId, entry.Frequency, entry.Timestamp)
		}

		statistics.Update(entry.Frequency)
		pool.Put(entry)
	}
}
