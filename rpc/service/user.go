package service

import (
	"gorm.io/gorm"
	"zg5/lmonth/message/user"
	"zg5/lmonth/rpc/model"
)

func GetUser(id int64) (*user.UserInfo, error) {
	newUser := model.NewUser()
	info, err := newUser.Get(id)
	if err != nil {
		return nil, err
	}
	return mysqlToPb(info)
}

func CreateUser(info *user.UserInfo) (*user.UserInfo, error) {
	newUser := model.NewUser()
	create, err := newUser.Create(pbToMysql(info))
	if err != nil {
		return nil, err
	}
	return mysqlToPb(create)
}

func UpdateUser(info *user.UserInfo) (*user.UserInfo, error) {
	newUser := model.NewUser()
	create, err := newUser.Update(pbToMysql(info))
	if err != nil {
		return nil, err
	}
	return mysqlToPb(create)
}

func DeleteUser(id int64) error {
	newUser := model.NewUser()
	err := newUser.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfoByTel(tel string) (*user.UserInfo, error) {
	newUser := model.NewUser()
	byTel, err := newUser.GetByTel(tel)
	if err != nil {
		return nil, err
	}
	return mysqlToPb(byTel)
}

func GetUserInfoByUsername(username string) (*user.UserInfo, error) {
	newUser := model.NewUser()
	byTel, err := newUser.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return mysqlToPb(byTel)
}

func mysqlToPb(user2 *model.User) (*user.UserInfo, error) {
	return &user.UserInfo{
		Id:       int64(user2.ID),
		Username: user2.Username,
		Password: user2.Password,
		Tel:      user2.Tel,
	}, nil
}

func pbToMysql(info *user.UserInfo) *model.User {
	return &model.User{
		Model: gorm.Model{
			ID: uint(info.Id),
		},
		Username: info.Username,
		Password: info.Password,
		Tel:      info.Tel,
	}
}
