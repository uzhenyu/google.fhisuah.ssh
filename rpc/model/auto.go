package model

import (
	"gorm.io/gorm"
	"zg5/lmonth/framework/mysql"
)

func AutoTable(fileName string) error {
	err := mysql.WithMysql(fileName, func(cli *gorm.DB) error {
		err := cli.AutoMigrate(new(User), new(Goods), new(Order))
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
