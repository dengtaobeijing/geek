package data

import (
	"context"
	"encoding/json"
	"network/app/sbi/admin/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.SbiRepo = (*sbiRepo)(nil)

// 实现 biz 定义的 SbiRepo 接口
type sbiRepo struct {
	data *Data
	log  *log.Helper
}

//PO
type Sbi struct {
	gorm.Model
	Id       int64
	Title    string
	Artist   string
	Price    float64
	CreateAt time.Time
}

func NewSbiRepo(data *Data, logger log.Logger) biz.SbiRepo {
	return &sbiRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/sbi")),
	}
}

func (r *sbiRepo) UpdateSbi(ctx context.Context, userId int64, b *biz.Sbi) error {
	o := Sbi{}
	result := r.data.db.WithContext(ctx).First(&o, b.Id)
	if result.Error != nil {
		return result.Error
	}

	o.Title = b.Title
	o.Artist = b.Artist
	o.UpdatedAt = time.Now()

	result = r.data.db.WithContext(ctx).Save(&o)
	originObj, _ := json.Marshal(o)
	newObj, _ := json.Marshal(b)
	err := r.data.db.Transaction(func(tx *gorm.DB) error {
		innerErr := tx.WithContext(ctx).Save(&o)
		if innerErr.Error != nil {
			return innerErr.Error
		}
		innerErr = insertIntegrating(ctx,
			string(originObj),
			string(newObj),
			userId, tx,
		)
		return innerErr.Error
	})
	return err
}
