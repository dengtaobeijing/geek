package data

import (
	"context"
	"time"

	"gorm.io/gorm"
)

//PO
// 积分表
type Operation struct {
	gorm.Model
	Id        int64
	UserId    int64
	OriginObj string
	NewObj    string
	CreateAt  time.Time
}

// 用例 ，原始值和新值 以对象json 形式存起来
func insertIntegrating(ctx context.Context, originObj, newObj string, userId int64, db *gorm.DB) *gorm.DB {
	o := Operation{
		UserId:    userId,
		OriginObj: originObj,
		NewObj:    newObj,
		CreateAt:  time.Now(),
	}
	return db.WithContext(ctx).Create(&o)
}
