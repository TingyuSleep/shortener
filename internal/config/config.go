package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct { // 匿名结构体
		DSN string
	}
	BaseString        string // base62指定的基础字符串
	ShortUrlBlackList []string
	ShortDomain       string
	CacheRedis        cache.CacheConf // redis缓存
}

type ShortUrlDB struct { // 命名结构体
	DSN string
}
