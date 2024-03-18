package service

import (
	"gorm.io/gorm"
	"zg5/lmonth/message/goods"
	"zg5/lmonth/rpc/model"
)

func GetGood(id int64) (*goods.GoodsInfo, error) {
	newGoods := model.NewGoods()
	get, err := newGoods.Get(id)
	if err != nil {
		return nil, err
	}
	return mysqlToPbG(get)
}

func GetGoods(offset, limit, goodsType int64) (infos []*goods.GoodsInfo, total int64, err error) {
	newGoods := model.NewGoods()
	getGoods, total, err := newGoods.GetGoods(offset, limit, goodsType)
	if err != nil {
		return nil, 0, err
	}
	for _, v := range getGoods {
		info, _ := mysqlToPbG(&v)
		infos = append(infos, info)
	}
	return
}

func GetGoodByName(name string) (*goods.GoodsInfo, error) {
	newGoods := model.NewGoods()
	byName, err := newGoods.GetByName(name)
	if err != nil {
		return nil, err
	}
	return mysqlToPbG(byName)
}

func CreateGoods(info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	newGoods := model.NewGoods()
	create, err := newGoods.Create(pbToMysqlG(info))
	if err != nil {
		return nil, err
	}
	return mysqlToPbG(create)
}

func UpdateGoods(info *goods.GoodsInfo) (*goods.GoodsInfo, error) {
	newGoods := model.NewGoods()
	create, err := newGoods.Update(pbToMysqlG(info))
	if err != nil {
		return nil, err
	}
	return mysqlToPbG(create)
}

func DeleteGoods(id int64) error {
	newGoods := model.NewGoods()
	err := newGoods.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStock(id, num int64) error {
	newGoods := model.NewGoods()
	err := newGoods.UpdateStock(id, num)
	if err != nil {
		return err
	}
	return nil
}

func mysqlToPbG(good *model.Goods) (*goods.GoodsInfo, error) {
	return &goods.GoodsInfo{
		Id:       int64(good.ID),
		Name:     good.Name,
		Price:    good.Price,
		VipPrice: good.VipPrice,
		Img:      good.Img,
		Stock:    good.Stock,
		Type:     good.Type,
		IdGoods:  good.IdGoods,
	}, nil
}

func pbToMysqlG(info *goods.GoodsInfo) *model.Goods {
	return &model.Goods{
		Model: gorm.Model{
			ID: uint(info.Id),
		},
		Name:     info.Name,
		Price:    info.Price,
		VipPrice: info.VipPrice,
		Img:      info.Img,
		Stock:    info.Stock,
		Type:     info.Type,
		IdGoods:  info.IdGoods,
	}
}
