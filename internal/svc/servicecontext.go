package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"
)

type ServiceContext struct {
	Config            config.Config
	ShortUrlMapModel  model.ShortUrlMapModel // 接口类型。 代表了 short_url_map表
	Sequence          sequence.Sequence
	ShortUrlBlackList map[string]struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	Conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	// 把配置文件中的黑名单加载到map，方便后续判断
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}

	return &ServiceContext{
		Config:           c,
		ShortUrlMapModel: model.NewShortUrlMapModel(Conn),
		Sequence:         sequence.NewMySQL(c.Sequence.DSN),
		//Sequence: 		sequence.NewRedis(RedisAddr),
		ShortUrlBlackList: m, // 短链接黑名单
	}
}
