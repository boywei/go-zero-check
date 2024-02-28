package middleware

import (
	"encoding/json"

	"github.com/boywei/go-zero-check/internal/model"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	RedisDb *redis.Client
)

func SetupRedisDb(addr, password string) error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	_, err := RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 查询key是否存在
func ExistKey(key string) bool {
	n, err := RedisDb.Exists(key).Result()
	if err != nil {
		log.Println("find exist model key error :", err)
	}
	if n == 0 {
		log.Println(key, "key no exist")
		return false
	}
	log.Println(key, "key exist")
	return true
}

// 根据Id获取模型缓存
func GetModelById(key string) (data *model.Uppaal, err error) {
	res, err := RedisDb.Get(key).Result()
	if err != nil {
		log.Println("GET redis error:", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(res), &data)
	return data, nil
}

// 根据文章Id设置缓存
func SetModelById(data *model.Uppaal, postid string) (err error) {
	// strdata, _ := json.Marshal(data)
	// key := uuid.NewRandom()
	// err = RedisDb.Set(key, strdata, 10*time.Second).Err()
	// // 设置key 10秒钟的过期时间
	// if err != nil {
	// 	log.Println("SET redis ERROR:", err)
	// 	return err
	// }
	return nil
}