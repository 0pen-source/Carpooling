package dao

import (
	"fmt"
	"time"

	"github.com/0pen-source/Carpooling/utils"
	"github.com/go-redis/redis"
)

var redisManager *Redis

func InitializeRedis() {
	options := utils.Must(redis.ParseURL(config.Redis.URL)).(*redis.Options)
	options.PoolSize = config.Redis.PoolSize
	redisManager = NewRedis(redis.NewClient(options))
}

// Redis _
type Redis struct {
	client *redis.Client
}

// NewRedis _
func NewRedis(client *redis.Client) *Redis {
	return &Redis{client: client}
}

// AdConsumption _
func (r *Redis) AdConsumption(adID string) int64 {
	key := fmt.Sprintf("consumption_%s", adID)
	consumption, err := r.client.Get(key).Int64()
	if err != nil && err != redis.Nil {
		return 0
	}

	return consumption
}

// Exceeds checks if frequency exceeds the limit of specific ad on specific device
func (r *Redis) Exceeds(limit int64, adID, uniqueDeviceID string) bool {
	key := r.key(adID, uniqueDeviceID)
	frequency, err := r.client.Get(key).Int64()
	if err != nil && err != redis.Nil {
		return false
	}

	return frequency >= limit
}

// Activate _
func (r *Redis) SetKey(key, value string) {

	r.client.Set(key, value, time.Hour*24*7)

}

// HasKey _
func (r *Redis) HasKey(key string) bool {
	_, err := r.client.Get(key).Result()

	if err != nil {
		return false
	}
	return true
}

// HasKey _
func (r *Redis) GetKey(key string) (string, error) {
	return r.client.Get(key).Result()

}

func (Redis) key(adID, uniqueDeviceID string) string {
	return fmt.Sprintf("impression_%s_%s", adID, uniqueDeviceID)
}

func (Redis) activatedKey(adID, uniqueDeviceID string) string {
	return fmt.Sprintf("activated_%s_%s", adID, uniqueDeviceID)
}
