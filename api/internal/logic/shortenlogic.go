package logic

import (
	"context"
	"github.com/lifezq/shorturl/rpc/transform/transformer"

	"github.com/lifezq/shorturl/api/internal/svc"
	"github.com/lifezq/shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (*types.ShortenResp, error) {
	// todo: add your logic here and delete this line

	// 手动代码开始
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return &types.ShortenResp{}, err
	}

	return &types.ShortenResp{
		Shorten: resp.Shorten,
	}, nil
	// 手动代码结束
}
