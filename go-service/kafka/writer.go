package ikafka

import (
	"crypto/tls"
	"strings"
	"time"

	"github.com/nc-minh/kafka-practices/config"
	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func NewWriter(config config.Config, topic string) *kafka.Writer {
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

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: strings.Split(config.KafkaBrokers, ","),
		Topic:   topic,
		Dialer:  dialer,
	})

	return w
}
