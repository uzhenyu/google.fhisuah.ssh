package client

import (
	"context"
	"zg5/lmonth/framework/grpc"
	"zg5/lmonth/message/order"
	"zg5/lmonth/rpc/conts"
)

type hande func(ctx context.Context, cli order.OrderClient) (interface{}, error)

func withOrder(ctx context.Context, han hande) (interface{}, error) {
	client, err := grpc.Client(conts.FilePath)
	if err != nil {
		return nil, err
	}
	orders := order.NewOrderClient(client)
	a, err := han(ctx, orders)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return a, nil
}

func Create(ctx context.Context, info *order.OrderInfo) (*order.OrderInfo, error) {
	i, err := withOrder(ctx, func(ctx context.Context, cli order.OrderClient) (interface{}, error) {
		createOrder, err := cli.CreateOrder(ctx, &order.CreateOrderRequest{Info: info})
		if err != nil {
			return nil, err
		}
		return createOrder.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*order.OrderInfo), nil
}
