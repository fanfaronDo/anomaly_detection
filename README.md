# Anomaly detection

Детектор аномалий — это инструмент, для выявления необычных отклонений в данных. Он может быть использован в различных областях, таких как финансовый анализ, системах мониторинга.
 
## Структура

-   Сервер. Содержет функцию генерации частот для определенного клиента.
-   Клиент. Сам детектор, вызавает функцию GenerateData на сервере, выявляет аномалии и складирует в базу postgres с определенным uuid сессии. 
- База данных postgres, таблица anomaly


## Установка

1. Установка среды. Установка преременных окружения.

    POSTGRES_USER=...
    POSTGRES_PASSWORD=...

2. Заппуск инф-структуры
    docker-compose up -d

3. На данный момент клиент запускается скриптом anomaly_detecor.sh в параметрах нужно указать коэфицент отклонения.
	bash anomaly_detector.sh 4

Скачивается образ с оф regestry и запускаетс скомпелированый бинарник в контейнере.  

Прокт выполнен в целях обучения. Но сам клиент можно применить выявления отклонений из своей программы.

Референс:
```bash
"github.com/fanfaronDo/anomaly_detection/pkg/client"
```

Пример:

```go
func Receiver(k float64, r *repo.Repository, stream pb.DataService_GenerateDataClient, wg *sync.WaitGroup) {
	statistics := &client.Statistics{}
	defer wg.Done()

	log.Printf("Reciver is running...\n")
	for {
		entry := pool.Get().(*pb.DataEntry)
		entry, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving data: %v", err)
		}

		if statistics.DetectAnomaly(entry.Frequency, k) {
			r.Create(entry)
			log.Printf("Received Data: Session ID: %s, Frequency: %f, Timestamp: %d\n", entry.SessionId, entry.Frequency, entry.Timestamp)
		}

		statistics.Update(entry.Frequency)
		pool.Put(entry)
	}
}

```
