package dao

import (
	"byteDemo/cmd/user/model"
	"byteDemo/kitex_gen/user"
	"log"
)

func CreatUser(req *user.Req) error {
	user := model.User{
		Id:   req.Id,
		Name: req.Name,
	}
	err := client.Create(&user).Error
	if err != nil {
		log.Println("mysql creat user failed")
		return err
	}
	return nil
}
