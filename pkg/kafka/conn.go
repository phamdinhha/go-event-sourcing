package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConn(ctx context.Context, kafkaCfg *Config) (*kafka.Conn, error) {
	return kafka.DialContext(ctx, "tcp", kafkaCfg.Brokers[0])
}
