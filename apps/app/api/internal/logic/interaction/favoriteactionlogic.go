package interaction

import (
	"TikTok/apps/app/api/apiVars"
	"TikTok/apps/interaction/rpc/interaction"
	"TikTok/apps/video/rpc/video"
	"context"
	"regexp"
	"strconv"

	"TikTok/apps/app/api/internal/svc"
	"TikTok/apps/app/api/internal/types"

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

	// 参数检查
	matched, err := regexp.MatchString("^\\d+$", strconv.FormatInt(req.VideoID, 10)) //是否为纯数字
	if strconv.FormatInt(req.VideoID, 10) == "" || matched == false {
		return &types.FavoriteActionResponse{
			RespStatus: types.RespStatus(apiVars.VideoIdRuleError),
		}, nil
	} else if (req.ActionType != 1) && (req.ActionType != 2) { //是否有除1或2的数字
		return &types.FavoriteActionResponse{
			RespStatus: types.RespStatus(apiVars.ActionTypeRuleError),
		}, nil
	}

	if req.Token == "" {
		return &types.FavoriteActionResponse{
			RespStatus: types.RespStatus(apiVars.NotLogged),
		}, nil
	}

	tokenID, err := l.svcCtx.JwtAuth.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.VideoRPC.Detail(l.ctx, &video.BasicVideoInfoReq{VideoId: req.VideoID})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.InteractionRPC.SendFavoriteAction(l.ctx, &interaction.FavoriteActionReq{
		UserId:     tokenID,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}

	return &types.FavoriteActionResponse{
		RespStatus: types.RespStatus(apiVars.Success),
	}, nil
}
