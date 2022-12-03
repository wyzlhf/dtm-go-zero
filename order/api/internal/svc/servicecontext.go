package svc

import (
	"dtm-go-zero/order/api/internal/config"
	"dtm-go-zero/order/rpc/orderclient"
	"dtm-go-zero/stock/rpc/stockclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	//此处使用orderclient
	OrderRpc orderclient.Order
	StockRpc stockclient.Stock
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		StockRpc: stockclient.NewStock(zrpc.MustNewClient(c.StockRpcConf)),
	}
}
