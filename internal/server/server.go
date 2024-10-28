package server

import (
	"fmt"
	"math/rand"
	"time"

	api "github.com/fanfaronDo/anomaly_detection/pkg/api/api/proto"
	"github.com/google/uuid"
)

type Server struct {
	api.UnimplementedDataServiceServer
}

func (serv *Server) GenerateData(stream api.DataService_GenerateDataServer) error {
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
