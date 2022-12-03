package logic

import (
	"context"
	"database/sql"
	"dtm-go-zero/stock/model"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"dtm-go-zero/stock/rpc/internal/svc"
	"dtm-go-zero/stock/rpc/types/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductLogic {
	return &DeductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductLogic) Deduct(in *stock.DeductReq) (*stock.DeductResp, error) {
	fmt.Printf("扣库存开始……")
	stock_,err:=l.svcCtx.StockModel.FindOneByGoodsId(l.ctx,in.GoodsId)
	if err!=nil && err!=model.ErrNotFound{
		return nil,status.Error(codes.Internal,err.Error())
	}
	if stock_==nil || stock_.Num<in.Num{
		return nil,status.Error(codes.Aborted,dtmcli.ResultFailure)
	}

	barrier,err:=dtmgrpc.BarrierFromGrpc(l.ctx)
	db,err:=sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err!=nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	if err:=barrier.CallWithDB(db, func(tx *sql.Tx) error {
		sqlResult,err:=l.svcCtx.StockModel.DeductStock(tx,in.GoodsId,in.Num)
		if err!=nil{
			return status.Error(codes.Internal,err.Error())
		}
		affected,err:=sqlResult.RowsAffected()
		if err!=nil{
			return status.Error(codes.Internal,err.Error())
		}
		if affected<=0{
			return status.Error(codes.Aborted,dtmcli.ResultFailure)
		}
		return nil
	});err!=nil{
		return nil, err
	}
	return &stock.DeductResp{}, nil
}
