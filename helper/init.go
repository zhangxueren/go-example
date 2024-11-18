package helper

import "context"

func Init(ctx context.Context) {
	initLogger(ctx)
	initRedis(ctx)
}
