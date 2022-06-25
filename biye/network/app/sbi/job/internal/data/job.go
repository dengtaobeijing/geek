package data

import (
	"context"
	"network/app/sbi/job/internal/biz"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

var _ biz.JobRepo = (*jobRepo)(nil)

// PO
type Integrating struct {
	gorm.Model
	Id       int64
	UserId   int64
	OrderId  int64
	Grade    int64
	CreateAt time.Time
}

type Orders struct {
	gorm.Model
	Id       int64
	UserId   int64
	SbiId    int64
	Price    float64
	Receiver string
	Address  string
	Mobile   string
}

// 实现 biz 定义的 SbiRepo 接口
type jobRepo struct {
	data *Data
	log  *log.Helper
}

func NewJobRepo(data *Data, logger log.Logger) biz.JobRepo {
	return &jobRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/job")),
	}
}

// 从 kafka 中捞取 订单,生成积分流水记录
func (r *jobRepo) IntegratingJob(ctx context.Context) error {

	partitionList, err := r.data.kc.Partitions("integrating") // topic
	if err != nil {
		panic(err)
	}
	root, _ := context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(root) // errgroup 的使用

	for partition := range partitionList {
		func(partition int32) {
			g.Go(func() error {
				pc, err := r.data.kc.ConsumePartition("integrating", partition, sarama.OffsetNewest)
				if err != nil {
					return err
				}
				msg := pc.Messages()
				json := string(msg.Value)
				var order Orders
				json.Unmarshal(&order, json) //这里暂时先忽略 反序列化错误
				err = insrtIntegrating(ctx, &order, r.data.db)
				if err != nil {
					return err
				}
				return nil
			})
		}(int32(partition))
	}

	return g.Wait()
}

func insrtIntegrating(ctx context.Context, order *Orders, db *gorm.DB) error {

	integrating := Integrating{
		UserId:   order.UserId,
		OrderId:  order.Id,
		Grade:    1, //每个订单新增一个积分，每个用户总积分，去 task 服务里计算
		CreateAt: time.Now(),
	}

	rst := db.WithContext(ctx).Create(&integrating)
	return rst.Error
}
