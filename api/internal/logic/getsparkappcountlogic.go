package logic

import (
	"context"
	"fmt"
	"time"
	"sparkAPPs/api/internal/svc"
	"sparkAPPs/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetSparkAppCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSparkAppCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSparkAppCountLogic {
	return GetSparkAppCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSparkAppCountLogic) GetSparkAppCount() ([]*types.AppCountReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("查询今天："+ time.Now().Format("20060102") +"和前10天的数据")
	
	sparkInfo, err := l.svcCtx.SparkAppsModel.FindOneByTime(time.Now().Format("20060102"))
	if err != nil {
		return nil, err
	}
	for k,v := range sparkInfo {
		fmt.Println(k,v.Id)
		fmt.Println(k,v.AppCounts)
		fmt.Println(k,v.CreateTime)
	}
	var result []*types.AppCountReply
	fmt.Println(result)
	for _,v := range sparkInfo {
		var reply types.AppCountReply
		reply.Id=v.Id
		reply.AppCounts=v.AppCounts
		reply.CreateTime=v.CreateTime
		result=append(result,&reply)
		// result[0]=&reply
	}
	// return &types.AppCountReply{
	// 	Id: sparkInfo.Id,
	// 	Count: sparkInfo.AppCounts,
	// 	CreateTime: sparkInfo.CreateTime,
	// }, nil
	return result,nil
}
