package interaction

import (
	"TikTok/apps/app/api/apiVars"
	"TikTok/apps/app/api/internal/middleware"
	"TikTok/apps/app/api/internal/svc"
	"TikTok/apps/app/api/internal/types"
	"TikTok/apps/interaction/rpc/interaction"
	"TikTok/apps/video/rpc/model"
	"TikTok/apps/video/rpc/video"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionRequest) (resp *types.FavoriteActionResponse, err error) {

	tokenID := l.ctx.Value(middleware.TokenIDKey).(int64)
	_, err = l.svcCtx.VideoRPC.Detail(l.ctx, &video.BasicVideoInfoReq{VideoId: req.VideoID})
	if errors.Is(err, model.ErrVideoNotFound) {
		return &types.FavoriteActionResponse{
			RespStatus: types.RespStatus(apiVars.VideoNotFound),
		}, nil
	}
	r, err := l.svcCtx.InteractionRPC.SendFavoriteAction(l.ctx, &interaction.FavoriteActionReq{
		UserId:     tokenID,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}
	if !r.IsSucceed {
		if req.ActionType == 1 {
			return &types.FavoriteActionResponse{
				RespStatus: types.RespStatus(apiVars.AlreadyLiked),
			}, nil
		} else {
			return &types.FavoriteActionResponse{
				RespStatus: types.RespStatus(apiVars.AlreadyUnLiked),
			}, nil
		}
	}
	return &types.FavoriteActionResponse{
		RespStatus: types.RespStatus(apiVars.Success),
	}, nil
}
