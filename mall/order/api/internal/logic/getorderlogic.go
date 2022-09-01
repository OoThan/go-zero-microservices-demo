package logic

import (
	"context"
	"errors"
	"fmt"
	"go-zero-microservices-demo/mall/user/rpc/userclient"

	"go-zero-microservices-demo/mall/order/api/internal/svc"
	"go-zero-microservices-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	if user.Name != "Hla" {
		return nil, errors.New("用户不存在")
	}
	return &types.OrderReply{
		Id:   req.Id,
		Name: "test Order",
	}, nil
}
