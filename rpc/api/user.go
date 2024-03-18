package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zg5/lmonth/message/user"
	"zg5/lmonth/rpc/service"
)

type ServiceUser struct {
	user.UnimplementedUserServer
}

func (s ServiceUser) GetUserByUsername(ctx context.Context, request *user.GetUserByUsernameRequest) (*user.GetUserByUsernameResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	username, err := service.GetUserInfoByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	return &user.GetUserByUsernameResponse{Info: username}, nil
}

func (s ServiceUser) GetUserByTel(ctx context.Context, request *user.GetUserByTelRequest) (*user.GetUserByTelResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Tel == "" {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	tel, err := service.GetUserInfoByTel(request.Tel)
	if err != nil {
		return nil, err
	}
	return &user.GetUserByTelResponse{Info: tel}, nil
}

func (s ServiceUser) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "不能为0")
	}
	getUser, err := service.GetUser(request.Id)
	if err != nil {
		return nil, err
	}
	return &user.GetUserResponse{Info: getUser}, nil
}

func (s ServiceUser) CreateUsers(ctx context.Context, request *user.CreateUsersRequest) (*user.CreateUsersResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Info == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Info.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "密码不能为空")
	}
	if request.Info.Tel == "" {
		return nil, status.Error(codes.InvalidArgument, "手机号不能为空")
	}
	createUser, err := service.CreateUser(request.Info)
	if err != nil {
		return nil, err
	}
	return &user.CreateUsersResponse{Info: createUser}, nil
}

func (s ServiceUser) DeleteUsers(ctx context.Context, request *user.DeleteUsersRequest) (*user.DeleteUsersResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	err := service.DeleteUser(request.Id)
	if err != nil {
		return nil, err
	}
	return &user.DeleteUsersResponse{}, nil
}

func (s ServiceUser) UpdateUsers(ctx context.Context, request *user.UpdateUsersRequest) (*user.UpdateUsersResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Info == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Info.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "不能为0")
	}
	updateUser, err := service.UpdateUser(request.Info)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUsersResponse{Info: updateUser}, nil
}
