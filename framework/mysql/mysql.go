package mysql

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zg5/lmonth/framework/nacos"
)

func WithMysql(fileName string, hand func(cli *gorm.DB) error) error {
	var cnf struct {
		Mysql struct {
			Username string `json:"Username"`
			Password string `json:"Password"`
			Host     string `json:"Host"`
			Port     string `json:"Port"`
			Database string `json:"Database"`
		} `json:"Mysql"`
	}
	config, err := nacos.GetConfig(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(config), &cnf)
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.Mysql.Username,
		cnf.Mysql.Password,
		cnf.Mysql.Host,
		cnf.Mysql.Port,
		cnf.Mysql.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, _ := db.DB()
	defer d.Close()
	if err = hand(db); err != nil {
		return err
	}
	return nil
}
