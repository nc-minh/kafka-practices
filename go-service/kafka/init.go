package ikafka

import (
	"crypto/tls"
	"time"

	"github.com/nc-minh/kafka-practices/config"
	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func NewKafkaDialer(config config.Config) *kafka.Dialer {
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

	return dialer
}
