package taskq

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

func NewClient(opt asynq.RedisClientOpt) *asynq.Client {

	client := asynq.NewClient(opt)
	return client
}

func NewServer(redis *redis.Client) *asynq.Server {
	srv := asynq.NewServer()

	return nil
}
