package interaction

import (
	"TikTok/apps/app/api/apiVars"
	"net/http"

	"TikTok/apps/app/api/internal/logic/interaction"
	"TikTok/apps/app/api/internal/svc"
	"TikTok/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := interaction.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJsonCtx(r.Context(), w, types.RespStatus(apiVars.InternalError))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
