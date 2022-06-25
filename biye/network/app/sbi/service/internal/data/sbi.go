package data

import (
	"context"
	"encoding/json"
	"fmt"
	"network/app/sbi/service/internal/biz"
	"network/pkg/util/pagination"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

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

func (r *sbiRepo) GetSbiById(ctx context.Context, id int64) (*biz.Sbi, error) {
	var sbi *biz.Sbi
	rSbi, err := r.getSbiFromRedis(ctx, id)
	if err != nil {
		return nil, err
	}
	sbi = rSbi
	if sbi != nil {
		return sbi, nil
	}
	sbi, err = r.getSbiFromDB(ctx, id)
	if err != nil {
		return nil, errors.NotFound("ALBUM_NOT_FOUND", fmt.Sprintf("sbi:%d not found", id))
	}
	err = r.setSbiToRedis(ctx, sbi)
	if err != nil {
		return nil, err
	}
	return sbi, nil
}

func (r *sbiRepo) getSbiFromDB(ctx context.Context, id int64) (*biz.Sbi, error) {
	a := Sbi{}
	result := r.data.db.WithContext(ctx).First(&a, id)
	return &biz.Sbi{
		Id:       a.Id,
		Title:    a.Title,
		Artist:   a.Artist,
		Price:    a.Price,
		CreateAt: a.CreateAt,
	}, result.Error
}

func (r *sbiRepo) getSbiFromRedis(ctx context.Context, id int64) (*biz.Sbi, error) {
	key := "getSbiFromRedis" + strconv.Itoa((int)(id))
	val, err := r.data.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var rst biz.Sbi
	json.Unmarshal([]byte(val), &rst)
	return &rst, nil
}

func (r *sbiRepo) setSbiToRedis(ctx context.Context, sbi *biz.Sbi) error {
	if val, err := json.Marshal(sbi); err != nil {
		return err
	} else {
		key := "getSbiFromRedis" + strconv.Itoa((int)(sbi.Id))
		err = r.data.rdb.Set(ctx, key, string(val), 2*time.Hour).Err()
		return err
	}
}

func (r *sbiRepo) ListSbi(ctx context.Context, pageNum, pageSize int64) (biz.ListSbi, error) {
	var as []Sbi
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&as)
	if result.Error != nil {
		return biz.ListSbi{}, result.Error
	}
	rv := make([]*biz.Sbi, 0)
	for _, o := range as {
		rv = append(rv, &biz.Sbi{
			Id:       o.Id,
			Title:    o.Title,
			Artist:   o.Artist,
			Price:    o.Price,
			CreateAt: o.CreateAt,
		})
	}
	var cnt int64
	r.data.db.WithContext(ctx).Model(&Sbi{}).Count(&cnt)
	return biz.ListSbi{
		Sbis:  rv,
		Count: cnt,
	}, nil
}

func (r *sbiRepo) CreateOrders(ctx context.Context, orders *biz.Orders) error {
	err := createOrders(ctx, orders, r.data.db)
	if err != nil {
		return err
	}

	b, err := json.Marshal(orders)
	if err != nil {
		return err
	}
	r.data.kp.Input() <- &sarama.ProducerMessage{
		Topic: "integrating", // 发送订单 ，供 积分job 消费
		Value: sarama.ByteEncoder(b),
	}
	return nil
}
