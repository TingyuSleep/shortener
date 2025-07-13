package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/config"
	"shortener/model"
)

type ServiceContext struct {
	Config           config.Config
	ShortUrlMapModel model.ShortUrlMapModel // 接口类型。 代表了 short_url_map表
}

func NewServiceContext(c config.Config) *ServiceContext {
	Conn := sqlx.NewMysql(c.ShortUrlDB.DSN)

	return &ServiceContext{
		Config:           c,
		ShortUrlMapModel: model.NewShortUrlMapModel(Conn),
	}
}
