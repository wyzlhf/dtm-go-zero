package logic

import (
	"context"
	"database/sql"
	"dtm-go-zero/order/model"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"dtm-go-zero/order/rpc/internal/svc"
	"dtm-go-zero/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateReq) (*order.CreateResp, error) {
	fmt.Printf("创建订单 in : %+v \n", in)
	barrier,err:=dtmgrpc.BarrierFromGrpc(l.ctx)
	db,err:=sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err!=nil{
		return nil, status.Error(codes.Internal,err.Error())
	}
	if err:=barrier.CallWithDB(db, func(tx *sql.Tx) error {
		order_:=new(model.Order)
		order_.GoodsId=in.GoodsId
		order_.Num=in.Num
		order_.UserId=in.UserId

		_,err=l.svcCtx.OrderModel.Insert(tx,order_)
		if err!=nil{
			return fmt.Errorf("创建订单失败 err : %v , order:%+v \n", err, order_)
		}
		return nil
	});err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}

	return &order.CreateResp{}, nil
}
