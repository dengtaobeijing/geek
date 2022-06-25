package data

import (
	"context"
	"encoding/json"
	"network/app/sbi/admin/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.OrdersRepo = (*ordersRepo)(nil)

// 实现 biz 定义的 OrdersRepo 接口
type ordersRepo struct {
	data *Data
	log  *log.Helper
}

//PO
type Orders struct {
	gorm.Model
	Id       int64
	UserId   int64
	SbiId    int64
	Price    float64
	Receiver string
	Address  string
	Mobile   string
	CreateAt time.Time
}

func NewOrdersRepo(data *Data, logger log.Logger) biz.OrdersRepo {
	return &ordersRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/orders")),
	}
}

func (r *ordersRepo) UpdateOrders(ctx context.Context, userId int64, b *biz.Orders) error {
	o := Orders{}
	result := r.data.db.WithContext(ctx).First(&o, b.Id)
	if result.Error != nil {
		return result.Error
	}

	o.Receiver = b.Receiver
	o.Address = b.Address
	o.Mobile = b.Address
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
