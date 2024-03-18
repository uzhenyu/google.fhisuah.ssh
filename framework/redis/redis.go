package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"zg5/lmonth/framework/nacos"
)

func WithRedis(fileName string, hand func(cli *redis.Client) error) error {
	var T struct {
		Redis struct {
			Host string `json:"Host"`
			Port string `json:"Port"`
		} `json:"Redis"`
	}
	config, err := nacos.GetConfig(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(config), &T)
	if err != nil {
		return err
	}
	re := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", T.Redis.Host, T.Redis.Port)})
	defer re.Close()
	if err = hand(re); err != nil {
		return err
	}
	return nil
}

func GetKey(ctx context.Context,fileName,key string) (string,error) {
	var data string
	err := WithRedis(fileName, func(cli *redis.Client) error {
		val, err := cli.Exists(ctx, key).Result()
		if err != nil {
			return nil
		}
		if val != 1 {
			return nil
		}
		data, err = cli.Get(ctx, key).Result()
		if err != nil{
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return data,nil
}

func ExistKey(ctx context.Context,fileName,key string) (bool,error) {
	//var data string
	err := WithRedis(fileName, func(cli *redis.Client) error {
		data, err := cli.Exists(ctx, key).Result()
		if err != nil {
			return nil
		}
		if data > 0 {
			return nil
		}
		return nil
	})
	if err != nil {
		return true, err
	}
	return false,nil
}
