package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct { // 匿名结构体
		DSN string
	}
}

type ShortUrlDB struct { // 命名结构体
	DSN string
}
