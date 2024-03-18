package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"zg5/lmonth/api/conts"
	"zg5/lmonth/framework/grpc"
	"zg5/lmonth/message/user"
)

type hand func(ctx context.Context, cli user.UserClient) (interface{}, error)

func withUser(ctx context.Context, han hand) (interface{}, error) {
	client, err := grpc.Client(conts.FileName)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	res, err := han(ctx, user.NewUserClient(client))
	logs.Info(err, 100)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUser(ctx context.Context, id int64) (*user.UserInfo, error) {
	i, err := withUser(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		getUser, err := cli.GetUser(ctx, &user.GetUserRequest{Id: id})
		if err != nil {
			return nil, err
		}
		return getUser.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*user.UserInfo), nil
}

func DeleteUsers(ctx context.Context, id int64) error {
	_, err := withUser(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		_, err := cli.DeleteUsers(ctx, &user.DeleteUsersRequest{Id: id})
		if err != nil {
			return nil, err
		}
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUsers(ctx context.Context, info *user.UserInfo) (*user.UserInfo, error) {
	i, err := withUser(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		users, err := cli.UpdateUsers(ctx, &user.UpdateUsersRequest{Info: info})
		if err != nil {
			return nil, err
		}
		return users.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*user.UserInfo), nil
}

func GetUserByUsername(ctx context.Context, username string) (*user.UserInfo, error) {
	i, err := withUser(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		usernames, err := cli.GetUserByUsername(ctx, &user.GetUserByUsernameRequest{Username: username})
		if err != nil {
			return nil, err
		}
		return usernames.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*user.UserInfo), nil
}

func GetUserByTel(ctx context.Context, tel string) (*user.UserInfo, error) {
	i, _ := withUser(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		byTel, err := cli.GetUserByTel(ctx, &user.GetUserByTelRequest{Tel: tel})
		logs.Info(byTel, 111)
		logs.Info(err, 1010)
		if err != nil {
			return nil, err
		}
		return byTel.Info, nil
	})
	if i == nil {
		logs.Info(i)
		return nil, errors.New("received nil UserInfo")
	}

	fmt.Printf("Type of i: %T\n", i) // 调试输出，查看 i 的实际类型

	userInfo, ok := i.(*user.UserInfo)
	if !ok {
		return nil, errors.New("failed to assert UserInfo")
	}

	return userInfo, nil
}

func CreateUsers(ctx context.Context, info *user.UserInfo) (*user.UserInfo, error) {
	i, err := withUser(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		users, err := cli.CreateUsers(ctx, &user.CreateUsersRequest{Info: info})
		if err != nil {
			return nil, err
		}
		return users.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*user.UserInfo), nil
}
