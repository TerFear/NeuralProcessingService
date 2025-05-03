package prc

import (
	"context"
	"go.uber.org/zap"
)

type Processing struct {
	log *zap.Logger
}

type PutInKafka interface {
	PutKafka(ctx context.Context, data string, DataType string) (res string, err error)
}
