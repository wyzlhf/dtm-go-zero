// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package stockclient

import (
	"context"

	"dtm-go-zero/stock/rpc/types/stock"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeductReq  = stock.DeductReq
	DeductResp = stock.DeductResp

	Stock interface {
		Deduct(ctx context.Context, in *DeductReq, opts ...grpc.CallOption) (*DeductResp, error)
		DeductRollback(ctx context.Context, in *DeductReq, opts ...grpc.CallOption) (*DeductResp, error)
	}

	defaultStock struct {
		cli zrpc.Client
	}
)

func NewStock(cli zrpc.Client) Stock {
	return &defaultStock{
		cli: cli,
	}
}

func (m *defaultStock) Deduct(ctx context.Context, in *DeductReq, opts ...grpc.CallOption) (*DeductResp, error) {
	client := stock.NewStockClient(m.cli.Conn())
	return client.Deduct(ctx, in, opts...)
}

func (m *defaultStock) DeductRollback(ctx context.Context, in *DeductReq, opts ...grpc.CallOption) (*DeductResp, error) {
	client := stock.NewStockClient(m.cli.Conn())
	return client.DeductRollback(ctx, in, opts...)
}
