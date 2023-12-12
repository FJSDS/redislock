package redislock

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

var redisClient = redis.NewClient(&redis.Options{
	Network: "tcp",
	Addr:    "127.0.0.1:6379", DB: 0,
	Password: "",
})

func TestClient_ObtainMany(t *testing.T) {
	r := require.New(t)
	c := New(redisClient)

	keys := []string{"123", "234"}
	lock, err := c.ObtainMany(context.Background(), 10, &Options{
		RetryStrategy: ExponentialBackoff(10*time.Millisecond, 512*time.Millisecond),
	}, keys...)
	r.NoError(err)
	defer lock.Release(context.Background())
	_, err = c.ObtainMany(context.Background(), 1, &Options{
		RetryStrategy: ExponentialBackoff(10*time.Millisecond, 512*time.Millisecond),
	}, keys...)
	r.Error(err)
	_, err = c.ObtainMany(context.Background(), 1, &Options{
		RetryStrategy: ExponentialBackoff(10*time.Millisecond, 512*time.Millisecond),
	}, keys[0], "2321")
	r.Error(err)
	_, err = c.ObtainMany(context.Background(), 1, &Options{
		RetryStrategy: ExponentialBackoff(10*time.Millisecond, 512*time.Millisecond),
	}, "31244", "2321")
	r.NoError(err)

}
