package data

import (
	"context"
	v1 "network/api/sbi/interface/v1"
	"network/app/sbi/interface/internal/biz"

	serviceV1 "network/api/sbi/service/v1"
	taskV1 "network/api/sbi/task/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.InterfaceRepo = (*interfaceRepo)(nil)

// 实现 biz 定义的 SbiRepo 接口
type interfaceRepo struct {
	data *Data
	log  *log.Helper
}

func NewInterfaceRepo(data *Data, logger log.Logger) biz.InterfaceRepo {
	return &interfaceRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/interface")),
	}
}

func (r *interfaceRepo) GetOrdersAndIntegrating(ctx context.Context, req *v1.GetOrdersAndIntegratingReq) (*v1.GetOrdersAndIntegratingReply, error) {
	// ps 这个获取订单的 接口是未真正实现了，这里偷懒了一下。哈哈
	orders, err := r.data.client.sbiClient.GetOrdersByUserId(ctx,
		&serviceV1.GetOrdersByUserIdReq{
			UserId: req.UserId,
		},
	)
	if err != nil {
		return nil, err
	}
	integrating, err := r.data.client.taskClient.GetIntegrating(ctx,
		&taskV1.GetIntegratingReq{
			UserId: req.UserId,
		},
	)
	if err != nil {
		return nil, err
	}
	// 拼装数据
	ordersReply := make([]*v1.GetOrdersAndIntegratingReply_Orders, 0)
	for _, order := range orders.Orders {
		ordersReply = append(ordersReply, &v1.GetOrdersAndIntegratingReply_Orders{
			UserId:   order.UserId,
			SbiId:    order.SbiId,
			Price:    order.Price,
			Receiver: order.Receiver,
			Address:  order.Address,
			Mobile:   order.Mobile,
		})
	}

	return &v1.GetOrdersAndIntegratingReply{
		UserId: req.UserId,
		Grade:  integrating.Grade,
		Orders: ordersReply,
	}, nil
}
