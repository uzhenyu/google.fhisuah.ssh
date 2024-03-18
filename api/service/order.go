package service

import (
	"context"
	"fmt"
	"zg5/lmonth/api/client"
	"zg5/lmonth/message/order"
)

func CreateOrder(ctx context.Context, info *order.OrderInfo) (*order.OrderInfo, error) {
	good, _ := client.GetGood(ctx, info.GoodsId)
	if good == nil {
		return nil, fmt.Errorf("该商品不存在")
	}
	//扣减库存
	err := client.UpdateStock(ctx, info.GoodsId, info.Num)
	if err != nil {
		return nil, err
	}
	create, err := client.Create(ctx, info)
	if err != nil {
		return nil, err
	}
	return create, nil
}
