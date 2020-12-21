package logic

import (
	"context"
	"time"
	"fmt"
	"sparkAPPs/api/internal/svc"
	"sparkAPPs/api/internal/types"
	"sparkAPPs/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type InsertSparkAppCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertSparkAppCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) InsertSparkAppCountLogic {
	return InsertSparkAppCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertSparkAppCountLogic) InsertSparkAppCount(req types.AppCountInsert) error {
	// todo: add your logic here and delete this line
	fmt.Println("插入数据")
	l.svcCtx.SparkAppsModel.Insert(model.SparkApps{
		AppCounts:       req.Count,
		CreateTime:      time.Now().Format("20060102"),
		// CreateTimestamp:  time.Now(),
	})
	return nil
}
