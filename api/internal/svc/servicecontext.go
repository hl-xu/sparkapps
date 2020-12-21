package svc

import (
	"sparkAPPs/api/internal/config"
    "sparkAPPs/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"

)

type ServiceContext struct {
	Config config.Config
	SparkAppsModel model.SparkAppsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	sm := model.NewSparkAppsModel(conn)

	return &ServiceContext{
		Config: c,
		SparkAppsModel: sm,
	}
}
