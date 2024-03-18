package model

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
	"zg5/lmonth/framework/mysql"
	"zg5/lmonth/framework/redis"
	"zg5/lmonth/rpc/conts"
)

type Goods struct {
	gorm.Model
	Name     string
	Price    string
	VipPrice string
	Img      string
	Stock    int64
	Type     int64
	IdGoods  int64
}

func NewGoods() *Goods {
	return new(Goods)
}

func (g *Goods) Get(id int64) (info *Goods, err error) {
	info = &Goods{}
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Model(g).Where("id = ?", id).First(info).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (g *Goods) GetGoods(offset, limit, goodsType int64) ([]Goods, int64, error) {
	var infos []Goods
	var total int64
	err := mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err := cli.Model(g).Where("type = ?", goodsType).Offset(int(offset)).Limit(int(limit)).Find(&infos).Error
		cli.Model(g).Count(&total)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos, total, nil
}

func (g *Goods) GetByName(name string) (info *Goods, err error) {
	info = &Goods{}
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Model(g).Where("name = ?", name).First(info).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (g *Goods) Create(goods *Goods) (info *Goods, err error) {
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Create(goods).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (g *Goods) Update(in *Goods) (info *Goods, err error) {
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Model(g).Where("id = ?", in.ID).Updates(in).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (g *Goods) Delete(id int64) error {
	err := mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err := cli.Model(g).Where("id = ?", id).Delete(g).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *Goods) UpdateStock(id, num int64) error {
	err := redis.WithRedis(conts.FilePath, func(cli *redis2.Client) error {
		err := cli.SetEX(context.Background(), "update_stock", 1, time.Second*1).Err()
		if err != nil {
			return err
		}
		defer cli.Del(context.Background(), "update_stock")
		return nil
	})
	if err != nil {
		return err
	}
	gets, _ := g.Get(id)
	gets.Stock -= num
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Model(g).Where("id = ?", id).Updates(gets).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
