package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
	Filter            *bloom.Filter // 布隆过滤器
}

func NewServiceContext(c config.Config) *ServiceContext {
	Conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	// 把配置文件中的黑名单加载到map，方便后续判断
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}

	// 初始化 redisBitset
	store, err := redis.NewRedis(c.CacheRedis[0].RedisConf, func(r *redis.Redis) {
		r.Type = redis.NodeType // 节点类型
	})
	if err != nil {
		logx.Error("Failed to initialize Redis: %v", err)
		panic(err)
	}

	// 实例化一个过滤器
	filter := bloom.New(store, "bloom_filter", 20*1<<20)

	return &ServiceContext{
		Config:           c,
		ShortUrlMapModel: model.NewShortUrlMapModel(Conn, c.CacheRedis),
		Sequence:         sequence.NewMySQL(c.Sequence.DSN),
		//Sequence: 		sequence.NewRedis(RedisAddr),
		ShortUrlBlackList: m, // 短链接黑名单
		Filter:            filter,
	}
}
