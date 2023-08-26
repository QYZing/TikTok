package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"TikTok/apps/user/rpc/internal/svc"
	"TikTok/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAvatarLogic {
	return &SetAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetAvatarLogic) SetAvatar(in *user.SetAvatarReq) (*user.SetAvatarResp, error) {
	isSucceed := true
	err := l.svcCtx.UserModel.UpdateByUserId(l.ctx, strconv.FormatInt(in.UserId, 10), "avatar", in.Url)
	if err != nil {
		isSucceed = false
	}

	//callback
	mqMap := make(map[string][]string, 10)
	mqMap[strconv.FormatInt(in.UserId, 10)] = []string{"avatar", strconv.FormatBool(isSucceed)}
	marshal, _ := json.Marshal(mqMap)
	callbackJSON := string(marshal)
	if err := l.svcCtx.KqPusherClient.Push(callbackJSON); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
		isSucceed = false
	}

	return &user.SetAvatarResp{
		IsSucceed: isSucceed,
	}, err
}
