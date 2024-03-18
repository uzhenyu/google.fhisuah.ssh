package model

import (
	"gorm.io/gorm"
	"zg5/lmonth/framework/mysql"
	"zg5/lmonth/rpc/conts"
)

type Order struct {
	gorm.Model
	GoodsId     int64
	Num         int64
	Total       string
	OrderIdCard string
}

func NewOrder() *Order {
	return new(Order)
}

func (o *Order) Create(order *Order) (info *Order, err error) {
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Create(order).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}
