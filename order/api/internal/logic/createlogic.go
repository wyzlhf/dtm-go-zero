package logic

import (
	"context"
	"dtm-go-zero/order/rpc/types/order"
	"dtm-go-zero/stock/rpc/types/stock"
	"fmt"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmgrpc"

	"dtm-go-zero/order/api/internal/svc"
	"dtm-go-zero/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
var dtmServer = "etcd://localhost:2379/dtmservice"
func (l *CreateLogic) Create(req *types.QuickCreateReq) (resp *types.QuickCreateResp, err error) {
	orderRpcBusiServer,err:=l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err!=nil{
		return nil, fmt.Errorf("下单异常超时")
	}
	stockRpcBusiServer,err:=l.svcCtx.Config.StockRpcConf.BuildTarget()
	if err!=nil{
		return nil, fmt.Errorf("下单异常超时")
	}
	createOrderReq:=&order.CreateReq{
		UserId: req.UserId,
		GoodsId: req.GoodsId,
		Num: req.Num,
	}
	deductReq:=&stock.DeductReq{
		GoodsId: req.GoodsId,
		Num: req.Num,
	}
	gid:=dtmgrpc.MustGenGid(dtmServer)
	saga:=dtmgrpc.NewSagaGrpc(dtmServer,gid).
		Add(orderRpcBusiServer+"/order.order/create",orderRpcBusiServer+"/order.order/createRollback",createOrderReq).
		Add(stockRpcBusiServer+"/stock.stock/deduct",stockRpcBusiServer+"/stock.stock/deductRollback",deductReq)
	err=saga.Submit()
	dtmimp.FatalIfError(err)
	if err!=nil{
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}

	return &types.QuickCreateResp{},nil
}
