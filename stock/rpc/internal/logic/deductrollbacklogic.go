package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"dtm-go-zero/stock/rpc/internal/svc"
	"dtm-go-zero/stock/rpc/types/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductRollbackLogic {
	return &DeductRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductRollbackLogic) DeductRollback(in *stock.DeductReq) (*stock.DeductResp, error) {
	fmt.Sprintf("库存回滚 in : %+v \n", in)
	barrier,err:=dtmgrpc.BarrierFromGrpc(l.ctx)
	db,err:=sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err!=nil{
		return nil, status.Error(codes.Internal,err.Error())
	}
	if err:=barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if err:=l.svcCtx.StockModel.AddStock(tx,in.GoodsId,in.Num);err!=nil{
			return fmt.Errorf("回滚库存失败 err : %v ,goodsId:%d , num :%d", err, in.GoodsId, in.Num)
		}
		return nil
	});err!=nil{
		logx.Errorf("err : %v \n", err)
		return nil, status.Error(codes.Internal,err.Error())
	}

	return &stock.DeductResp{}, nil
}
