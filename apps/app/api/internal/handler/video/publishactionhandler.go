package video

import (
	"TikTok/apps/app/api/apiVars"
	"net/http"

	"TikTok/apps/app/api/internal/logic/video"
	"TikTok/apps/app/api/internal/svc"
	"TikTok/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewPublishActionLogic(r.Context(), svcCtx)
		resp, err := l.PublishAction(&req, r)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJsonCtx(r.Context(), w, types.RespStatus(apiVars.InternalError))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
