package repository

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gomyred/config"
	"strconv"
	"time"
)

var (
	conRedis *redis.Client
)

// NewRedis
func NewRedis(c config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	conRedis = client

	return conRedis, nil
}

// GetConnectionRedis
func GetConnectionRedis() *redis.Client {
	return conRedis
}

// SetCache
func SetCache(page int, dataCache []byte, TTL int) {

	err := conRedis.Set(strconv.Itoa(page), dataCache, time.Second*time.Duration(TTL)).Err()
	if err != nil {
		panic(err)
	}
}

func GetCache(page int) (string, error) {
	val2, err := conRedis.Get(strconv.Itoa(page)).Result()
	if err != nil {
		return "", err
	}

	return val2, nil
}
