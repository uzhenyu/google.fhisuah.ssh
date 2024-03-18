package model

import (
	"gorm.io/gorm"
	"zg5/lmonth/framework/mysql"
	"zg5/lmonth/rpc/conts"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Tel      string
}

func NewUser() *User {
	return new(User)
}

func (u *User) Get(id int64) (info *User, err error) {
	info = new(User)
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Where("id = ?", id).First(info).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}

func (u *User) Create(in *User) (info *User, err error) {
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Create(in).Error
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

func (u *User) Update(in *User) (info *User, err error) {
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Model(u).Where("id = ?", in.ID).Updates(in).Error
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

func (u *User) Delete(id int64) error {
	err := mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err := cli.Where("id = ?", id).Delete(u).Error
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

func (u *User) GetByTel(tel string) (info *User, err error) {
	info = &User{}
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Where("tel = ?", tel).First(info).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func (u *User) GetByUsername(username string) (info *User, err error) {
	info = &User{}
	err = mysql.WithMysql(conts.FilePath, func(cli *gorm.DB) error {
		err = cli.Where("username = ?", username).First(info).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}
