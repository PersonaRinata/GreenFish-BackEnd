package initialize

import (
	"GreenFish/server/common/consts"
	"GreenFish/server/service/user/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GlobalServerConfig.RedisInfo.Host, config.GlobalServerConfig.RedisInfo.Port),
		Password: config.GlobalServerConfig.RedisInfo.Password,
		DB:       consts.RedisUserClientDB,
	})
	return client
}
