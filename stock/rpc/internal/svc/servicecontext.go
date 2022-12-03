package svc

import (
	"dtm-go-zero/stock/model"
	"dtm-go-zero/stock/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	StockModel model.StockModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		StockModel: model.NewStockModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
