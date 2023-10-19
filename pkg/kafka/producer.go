package kafka

import (
	"context"

	"github.com/phamdinhha/go-event-sourcing/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	log     logger.Logger
	brokers []string
	w       *kafka.Writer
}

func NewProducer(log logger.Logger, brokers []string) *producer {
	return &producer{
		log:     log,
		brokers: brokers,
		w:       NewWriter(brokers, kafka.LoggerFunc(log.Errorf)),
	}
}

func NewAsyncProducer(log logger.Logger, brokers []string) *producer {
	return &producer{
		log:     log,
		brokers: brokers,
		w:       NewAsyncWriter(brokers, kafka.LoggerFunc(log.Errorf), log),
	}
}

func NewAsyncProducerWithCallback(log logger.Logger, brokers []string, cb AsyncWriterCallback) *producer {
	return &producer{
		log:     log,
		brokers: brokers,
		w:       NewRequireNoneWriter(brokers, kafka.LoggerFunc(log.Errorf), log)}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	if err := p.w.WriteMessages(ctx, msgs...); err != nil {
		return err
	}
	return nil
}

func (p *producer) Close() error {
	return p.w.Close()
}
