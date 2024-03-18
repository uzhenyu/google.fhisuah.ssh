package service

import (
	"context"
	"fmt"
	"zg5/lmonth/api/client"
	"zg5/lmonth/message/goods"
)

func CreateGood(ctx context.Context, info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	good, _ := client.GetGood(ctx, info.Id)
	if good != nil {
		return nil, fmt.Errorf("该商品已存在")
	}
	createGood, err := client.CreateGood(ctx, info)
	if err != nil {
		return nil, err
	}
	return createGood, nil
}

func DeleteGood(ctx context.Context, id int64) error {
	err := client.DeleteGood(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateGood(ctx context.Context, info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	good, _ := client.GetGood(ctx, info.Id)
	if good == nil {
		return nil, fmt.Errorf("该商品不存在")
	}
	updateGood, err := client.UpdateGood(ctx, info)
	if err != nil {
		return nil, err
	}
	return updateGood, nil
}

func SelectGoods(ctx context.Context, id int64) (*goods.GoodsInfo, error) {
	good, _ := client.GetGood(ctx, id)
	if good == nil {
		return nil, fmt.Errorf("该商品不存在")
	}
	return good, nil
}

func GetGoods(ctx context.Context, offset, limit, goodsType int64) ([]*goods.GoodsInfo, error) {
	getGoods, err := client.GetGoods(ctx, offset, limit, goodsType)
	if err != nil {
		return nil, err
	}
	return getGoods, nil
}
