package prc

import (
	"context"
	"go.uber.org/zap"
)

type Processing struct {
	log          *zap.Logger
	CheckHistory CheckRedis
	Kafka        PutInKafka
}

func New(log *zap.Logger, check_redis CheckRedis, kafka PutInKafka) *Processing {
	return &Processing{
		log:          log,
		CheckHistory: check_redis,
		Kafka:        kafka,
	}
}

type CheckRedis interface {
	CheckHash(ctx context.Context, answer string) (data string, err error)
}

type PutInKafka interface {
	ProcessingML(ctx context.Context, userID string, data string, Topic string) (err error)
}

func (p *Processing) Process(ctx context.Context, userID string, data string, topic string) (res string, err error) {
	res, err = p.CheckHistory.CheckHash(ctx, data)

	switch res {

	case "":
		err = p.Kafka.ProcessingML(ctx, userID, data, topic)

		if err != nil {
			return "", err
		}

		return "Waiting answer", nil

	default:
		p.log.Debug("Check hash success")
		return res, nil
	}

}
