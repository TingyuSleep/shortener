package sequence

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// redis 实现取号器功能
type Redis struct {
	// redis 连接
	conn sqlx.SqlConn
}

func NewRedis(redisAddr string) Sequence {
	return &Redis{}
}

func (r *Redis) Next() (seq uint64, err error) {
	// TODO:使用redis实现发号器
	return
}
