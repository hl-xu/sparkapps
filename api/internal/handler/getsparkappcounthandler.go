package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"sparkAPPs/api/internal/logic"
	"sparkAPPs/api/internal/svc"
)

func GetSparkAppCountHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetSparkAppCountLogic(r.Context(), ctx)
		resp, err := l.GetSparkAppCount()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
