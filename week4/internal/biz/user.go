package biz

// 定义data实现

type UserDO struct {
	Id       string
	Name     string
	Age      int
	UserName string
}

type UserRepo interface {
	Get(string) UserDO
}

type UserBiz struct {
	repo UserRepo
}

func NewUserBiz(repo UserRepo) *UserBiz {
	return &UserBiz{repo: repo}
}

func (u *UserBiz) Get(id string) UserDO {
	return u.repo.Get(id)
}
