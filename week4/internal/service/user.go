package service

import (
	"geek/week4/api/test"
	"geek/week4/internal/biz"
	"log"
)

type UserService struct {
	bz  *biz.UserBiz
	log *log.Logger
}

func NewUserService(bz *biz.UserBiz, log *log.Logger) *UserService {
	return &UserService{bz: bz, log: log}
}

func (u *UserService) SayHello(id string) *test.HelloDTO {
	userDO := u.bz.Get(id)
	return &test.HelloDTO{
		Name:     userDO.Name,
		UserName: userDO.UserName,
	}
}
