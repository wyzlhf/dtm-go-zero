// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package orderclient

import (
	"context"

	"dtm-go-zero/order/rpc/types/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateReq  = order.CreateReq
	CreateResp = order.CreateResp

	Order interface {
		Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
		CreateRollback(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultOrder) CreateRollback(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateRollback(ctx, in, opts...)
}
