package ikafka

import (
	"crypto/tls"
	"strings"
	"time"

	"github.com/nc-minh/kafka-practices/config"
	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func NewReader(config config.Config, topic string) *kafka.Reader {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		SASLMechanism: plain.Mechanism{
			Username: config.KafkaApiKey,
			Password: config.KafkaApiSecret,
		},
	}

	batchSize := int(10e6) // 10MB

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(config.KafkaBrokers, ","),
		Topic:    topic,
		Dialer:   dialer,
		MinBytes: 10e3, // 10KB
		MaxBytes: batchSize,
		// GroupID:  "NCM-test",
	})

	return r
}
