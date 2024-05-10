package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nc-minh/kafka-practices/config"
	ikafka "github.com/nc-minh/kafka-practices/kafka"
	"github.com/segmentio/kafka-go"
)

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	router := gin.Default()
	// Load configuration
	config, _ := config.LoadConfig(".")

	// Create Kafka writer and reader
	w := ikafka.NewWriter(config, "order")
	// r := ikafka.NewReader(config, "order")

	var wg sync.WaitGroup
	wg.Add(2)

	defer w.Close()
	// defer r.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// Start a goroutine to read messages from Kafka
	// go func() {
	// 	defer wg.Done()
	// 	log.Println("Reading messages from Kafka...")
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		default:
	// 			m, err := r.ReadMessage(context.Background())
	// 			if err != nil {
	// 				log.Printf("failed to read message: %v", err)
	// 				continue
	// 			}
	// 			fmt.Printf("message at offset=%d, key=%s, value=%s, partition=%d\n", m.Offset, string(m.Key), string(m.Value), m.Partition)
	// 		}
	// 	}
	// }()

	// err = r.CommitMessages(ctx, m)
	// 		if err != nil {
	// 			log.Printf("failed to commit message: %v", err)
	// 			return
	// 		}

	// go func() {
	// 	err := w.WriteMessages(context.Background(),
	// 		kafka.Message{
	// 			Key:   []byte("XXX-MMM-SSS"),
	// 			Value: []byte("NCM"),
	// 		},
	// 	)
	// 	log.Println("sent ok")
	// 	if err != nil {
	// 		log.Println("failed to write messages:", err)
	// 	}

	// }()

	go func() {
		defer wg.Done()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("Shutting down...")
		// cancel()
	}()

	router.POST("/ping", func(c *gin.Context) {
		var data Data

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(data.Key),
				Value: []byte(data.Value),
			},
		)
		if err != nil {
			log.Println("failed to write messages:", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"key":     data.Key,
			"value":   data.Value,
		})
	})
	router.Run(":8080")

}
