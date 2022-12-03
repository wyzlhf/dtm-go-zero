package svc

import (
	"dtm-go-zero/order/model"
	"dtm-go-zero/order/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OrderModel: model.NewOrderModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
