package handler

import (
	"net/http"

	"sparkAPPs/api/internal/logic"
	"sparkAPPs/api/internal/svc"
	"sparkAPPs/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func InsertSparkAppCountHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppCountInsert
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewInsertSparkAppCountLogic(r.Context(), ctx)
		err := l.InsertSparkAppCount(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
