package main

import (
	"context"
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fanfaronDo/anomaly_detection/internal/client"
	pb "github.com/fanfaronDo/anomaly_detection/pkg/api/api/proto"
	"github.com/fanfaronDo/anomaly_detection/pkg/config"
	repo "github.com/fanfaronDo/anomaly_detection/pkg/repository"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type Connector interface {
	Connect(c *config.Config) (*gorm.DB, error)
}

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

	host := myConfig.ServerTransmitterHost + ":" + string(myConfig.ServerTransmitterPort)
	connectionToClient, err := grpc.Dial(host,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer connectionToClient.Close()
	connector := Connector{}
	conn, err := connector.Connect(myConfig)
	if err != nil {
		log.Fatalf("Errer connect to db %s\n", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Receiver(*flagK, *myConfig, logFile, wg)
	wg.Wait()
}

func Receiver(k float64, logFile *os.File, wg *sync.WaitGroup) {
	statistics := &client.Statistics{}
	defer wg.Done()

	log.Printf("Client connect to %s\n", host)

	conn, err := repo.NewConnector(cfg)
	if err != nil {
		log.Fatalf("Errer connect to db %s\n", err)
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
		}

		if statistics.DetectAnomaly(entry.Frequency, k) {
			r.Create(*entry)
			log.Printf("Received Data: Session ID: %s, Frequency: %f, Timestamp: %d\n", entry.SessionId, entry.Frequency, entry.Timestamp)
		}

		statistics.Update(entry.Frequency)
		pool.Put(entry)
	}
}
