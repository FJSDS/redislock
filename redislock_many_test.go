package redislock

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
)

var redisClient = redis.NewClient(&redis.Options{
	Network: "tcp",
	Addr:    "10.23.20.53:6379", DB: 9,
	Password: "iggcdl5,.",
})



func TestClient_ObtainMany(t *testing.T) {
	r:=require.New(t)
	c:=New(redisClient)



	keys:=[]string{"123","234"}
	lock,err:=c.ObtainMany(context.Background(),10,&Options{
		RetryStrategy: ExponentialBackoff(10,512),
	},keys...)
	r.NoError(err)
	defer lock.Release(context.Background())
	_,err = c.ObtainMany(context.Background(),1,&Options{
		RetryStrategy: ExponentialBackoff(10,512),
	},keys...)
	r.Error(err)
	_,err = c.ObtainMany(context.Background(),1,&Options{
		RetryStrategy: ExponentialBackoff(10,512),
	},keys[0],"2321")
	r.Error(err)
	_,err = c.ObtainMany(context.Background(),1,&Options{
		RetryStrategy: ExponentialBackoff(10,512),
	},"31244","2321")
	r.NoError(err)

}
