package data

import "geek/week4/internal/biz"

type UserPO struct {
	Id       string
	Name     string
	Age      int
	UserName string
	Password string
}
type userRepo struct {
}

func NewUserRepo() biz.UserRepo {
	return &userRepo{}
}

func (u *userRepo) Get(id string) biz.UserDO {
	user := dbFind(id)
	return biz.UserDO{
		Id:       user.Id,
		Name:     user.Name,
		Age:      user.Age,
		UserName: user.UserName,
	}
}

func dbFind(id string) UserPO {
	return UserPO{
		Id:       id,
		Name:     "Name" + id,
		Age:      20,
		UserName: "user_" + id,
		Password: "123456",
	}
}
