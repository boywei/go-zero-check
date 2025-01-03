package middleware

import (
	"encoding/json"
	"errors"
	"github.com/boywei/go-zero-check/internal/model"
	"github.com/boywei/go-zero-check/internal/util/global"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
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

// ExistKey 查询key是否存在
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

// GetModelById 根据Id获取模型缓存
func GetModelById(key string) (data *model.Uppaal, err error) {
	res, err := RedisDb.Get(key).Result()
	if err != nil {
		log.Errorln("GET redis error:", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Errorln("Get redis model error: ", err)
	}
	return data, nil
}

// SetModel (不再使用redis了, 所以以下代码除了返回uuid都废弃了)设置模型的缓存
func SetModel(data *model.Uppaal) (key string, err error) {
	strdata, err := json.Marshal(data)
	if err != nil {
		log.Errorln("Json Marshal error: ", err)
	}
	key = uuid.NewString()
	// 模型保存时间为1天
	err = RedisDb.Set(key, strdata, time.Hour*24).Err()
	if err != nil {
		log.Println("SET redis ERROR:", err)
		return "", err
	}
	return key, nil
}

// DeleteModelById 根据id删除对应的模型
func DeleteModelById(id string) (err error) {
	// 删除redis缓存
	//RedisDb.Del(id)
	// 删除modelMap中的内容
	m, ok := global.ModelIdMap[id]
	if !ok {
		log.Warnln(id, "not exists")
		return errors.New("id not exists")
	}
	m.RemoveDir()
	delete(global.ModelIdMap, id)
	return nil
}
