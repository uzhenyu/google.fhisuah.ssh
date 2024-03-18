package service

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	redis2 "github.com/go-redis/redis/v8"
	"time"
	"zg5/lmonth/api/client"
	"zg5/lmonth/api/conts"
	"zg5/lmonth/api/utils"
	"zg5/lmonth/framework/redis"
	"zg5/lmonth/message/user"
)

type LoginReq struct {
	UserId int64
	Token  string
}

func Login(ctx context.Context, tel, code string) (*LoginReq, error) {
	byTel, _ := client.GetUserByTel(ctx, tel)
	if byTel == nil {
		return nil, fmt.Errorf("该用户不存在")
	}
	err := redis.WithRedis(conts.FileName, func(cli *redis2.Client) error {
		result, _ := cli.Get(context.Background(), tel).Result()
		if code != result {
			return fmt.Errorf("验证码错误")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	token, err := utils.GetJwtToken(conts.Servername, time.Now().Unix(), int64(time.Minute*60), byTel.Id)
	if err != nil {
		return nil, err
	}
	return &LoginReq{
		UserId: byTel.Id,
		Token:  token,
	}, nil
}

func RegisterUser(ctx context.Context, tel, password, code string) (*LoginReq, error) {
	byTel, _ := client.GetUserByTel(ctx, tel)
	if byTel != nil {
		return nil, fmt.Errorf("该手机号已注册")
	}
	err := redis.WithRedis(conts.FileName, func(cli *redis2.Client) error {
		result, _ := cli.Get(context.Background(), tel).Result()
		if code != result {
			return fmt.Errorf("验证码错误")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	users, err := client.CreateUsers(ctx, &user.UserInfo{
		Password: password,
		Tel:      tel,
	})
	if err != nil {
		return nil, err
	}
	return &LoginReq{
		UserId: users.Id,
		Token:  "",
	}, nil
}

func SendMessage(ctx context.Context, tel string) error {
	logs.Info(tel)
	code, err := utils.Code(tel)
	if err != nil {
		return err
	}
	err = redis.WithRedis(conts.FileName, func(cli *redis2.Client) error {
		cli.Set(context.Background(), tel, code, time.Minute*3)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
