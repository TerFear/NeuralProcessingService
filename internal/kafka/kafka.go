package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaProduser struct {
	Kafka *kafka.Writer
	log   *zap.Logger
}

func New(log *zap.Logger) *KafkaProduser {

	k := &KafkaProduser{
		log: log,
		Kafka: &kafka.Writer{
			Addr:     kafka.TCP("localhost:9092"),
			Async:    true,
			Balancer: &kafka.LeastBytes{},
		},
	}

	return k
}

func (k *KafkaProduser) ProcessingML(ctx context.Context, userID string, data string, Topic string) (err error) {
	msg := &kafka.Message{
		Topic: Topic,
		Value: []byte(data),
		Key:   []byte(userID),
	}

	err = k.Kafka.WriteMessages(ctx, *msg)

	if err != nil {
		return err
	}

	return nil
}
