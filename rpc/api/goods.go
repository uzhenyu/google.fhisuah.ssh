package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zg5/lmonth/message/goods"
	"zg5/lmonth/rpc/service"
)

type ServiceGood struct {
	goods.UnimplementedGoodsServer
}

func (s ServiceGood) UpdateStock(ctx context.Context, request *goods.UpdateStockRequest) (*goods.UpdateStockResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "不能为0")
	}
	err := service.UpdateStock(request.Id, request.Num)
	if err != nil {
		return nil, err
	}
	return &goods.UpdateStockResponse{Info: nil}, nil
}

func (s ServiceGood) GetGood(ctx context.Context, request *goods.GetGoodRequest) (*goods.GetGoodResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "不能为0")
	}
	good, err := service.GetGood(request.Id)
	if err != nil {
		return nil, err
	}
	return &goods.GetGoodResponse{Info: good}, nil
}

func (s ServiceGood) GetGoodByName(ctx context.Context, request *goods.GetGoodsByNameRequest) (*goods.GetGoodsByNameResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	name, err := service.GetGoodByName(request.Name)
	if err != nil {
		return nil, err
	}
	return &goods.GetGoodsByNameResponse{Info: name}, nil
}

func (s ServiceGood) GetGoods(ctx context.Context, request *goods.GetGoodsRequest) (*goods.GetGoodsResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Limit == 0 || request.Offset == 0 || request.Type == 0 {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	getGoods, i, err := service.GetGoods(request.Offset, request.Limit, request.Type)
	if err != nil {
		return nil, err
	}
	return &goods.GetGoodsResponse{
		Info:  getGoods,
		Total: i,
	}, nil
}

func (s ServiceGood) CreateGood(ctx context.Context, request *goods.CreateGoodRequest) (*goods.CreateGoodResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	createGoods, err := service.CreateGoods(request.Info)
	if err != nil {
		return nil, err
	}
	return &goods.CreateGoodResponse{Info: createGoods}, nil
}

func (s ServiceGood) DeleteGood(ctx context.Context, request *goods.DeleteGoodRequest) (*goods.DeleteGoodResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	if request.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "不能为0")
	}
	err := service.DeleteGoods(request.Id)
	if err != nil {
		return nil, err
	}
	return &goods.DeleteGoodResponse{}, nil
}

func (s ServiceGood) UpdateGood(ctx context.Context, request *goods.UpdateGoodRequest) (*goods.UpdateGoodResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	updateGoods, err := service.UpdateGoods(request.Info)
	if err != nil {
		return nil, err
	}
	return &goods.UpdateGoodResponse{Info: updateGoods}, nil
}
