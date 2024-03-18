package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zg5/lmonth/message/order"
	"zg5/lmonth/rpc/service"
)

type ServerOrder struct {
	order.UnimplementedOrderServer
}

func (s ServerOrder) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "不能为空")
	}
	createOrder, err := service.CreateOrder(request.Info)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{Info: createOrder}, nil
}
