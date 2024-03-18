package service

import (
	"gorm.io/gorm"
	"zg5/lmonth/message/order"
	"zg5/lmonth/rpc/model"
)

func CreateOrder(info *order.OrderInfo) (*order.OrderInfo, error) {
	newOrder := model.NewOrder()
	create, err := newOrder.Create(pbToMy(info))
	if err != nil {
		return nil, err
	}
	return myToPb(create)
}

func myToPb(order2 *model.Order) (*order.OrderInfo, error) {
	return &order.OrderInfo{
		Id:          int64(order2.ID),
		GoodsId:     order2.GoodsId,
		Num:         order2.Num,
		Total:       order2.Total,
		OrderIdCard: order2.OrderIdCard,
	}, nil
}

func pbToMy(info *order.OrderInfo) *model.Order {
	return &model.Order{
		Model: gorm.Model{
			ID: uint(info.Id),
		},
		GoodsId:     info.GoodsId,
		Num:         info.Num,
		Total:       info.Total,
		OrderIdCard: info.OrderIdCard,
	}
}
