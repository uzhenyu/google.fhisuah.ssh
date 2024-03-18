package client

import (
	"context"
	"zg5/lmonth/framework/grpc"
	"zg5/lmonth/message/goods"
	"zg5/lmonth/rpc/conts"
)

type hands func(ctx context.Context, cli goods.GoodsClient) (interface{}, error)

func withGoods(ctx context.Context, han hands) (interface{}, error) {
	client, err := grpc.Client(conts.FilePath)
	if err != nil {
		return nil, err
	}
	newGoods := goods.NewGoodsClient(client)
	a, err := han(ctx, newGoods)
	defer client.Close()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func GetGood(ctx context.Context, id int64) (*goods.GoodsInfo, error) {
	i, err := withGoods(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		good, err := cli.GetGood(ctx, &goods.GetGoodRequest{Id: id})
		if err != nil {
			return nil, err
		}
		return good.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*goods.GoodsInfo), nil
}

func GetGoods(ctx context.Context, offset, limit, goodsType int64) ([]*goods.GoodsInfo, error) {
	i, err := withGoods(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		getGoods, err := cli.GetGoods(ctx, &goods.GetGoodsRequest{
			Offset: offset,
			Limit:  limit,
			Type:   goodsType,
		})
		if err != nil {
			return nil, err
		}
		return getGoods.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.([]*goods.GoodsInfo), nil
}

func GetGoodsByName(ctx context.Context, name string) (*goods.GoodsInfo, error) {
	i, err := withGoods(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		byName, err := cli.GetGoodByName(ctx, &goods.GetGoodsByNameRequest{Name: name})
		if err != nil {
			return nil, err
		}
		return byName.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*goods.GoodsInfo), nil
}

func CreateGood(ctx context.Context, info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	i, err := withGoods(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		good, err := cli.CreateGood(ctx, &goods.CreateGoodRequest{Info: info})
		if err != nil {
			return nil, err
		}
		return good.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*goods.GoodsInfo), nil
}

func DeleteGood(ctx context.Context, id int64) error {
	_, err := withGoods(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		good, err := cli.DeleteGood(ctx, &goods.DeleteGoodRequest{Id: id})
		if err != nil {
			return nil, err
		}
		return good.Info, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateGood(ctx context.Context, info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	i, err := withGoods(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		good, err := cli.UpdateGood(ctx, &goods.UpdateGoodRequest{Info: info})
		if err != nil {
			return nil, err
		}
		return good.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return i.(*goods.GoodsInfo), nil
}

func UpdateStock(ctx context.Context, id, num int64) error {
	_, err := withGoods(ctx, func(ctx context.Context, cli goods.GoodsClient) (interface{}, error) {
		stock, err := cli.UpdateStock(ctx, &goods.UpdateStockRequest{
			Id:  id,
			Num: num,
		})
		if err != nil {
			return nil, err
		}
		return stock.Info, nil
	})
	if err != nil {
		return err
	}
	return nil
}
