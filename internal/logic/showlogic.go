package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	Err404 = errors.New("404")
)

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 自己写缓存，surl -> lurl
// go-zero自带的缓存，surl -> 数据行

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 查看短链接，输入q1mi.cn/lucky ---> 重定向到真实的链接
	// req.ShortUrl = lucky
	// 1. 根据短链接查询原始的长链接
	// 1.0 布隆过滤器：不存在的短链接直接返回404，不需要后续处理
	// a. 基于内存版本,服务器重启之后就没了，多以每次重启都要重新加载一下已有的短链接（查数据库）
	// b. 基于redis版本,go-zero自带
	exist, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("Bloom Filter.Exists failed", logx.LogField{Key: "err", Value: err.Error()},
			logx.LogField{Key: "shortUrl", Value: req.ShortUrl})
	}
	if !exist {
		logx.Errorw("shortUrl not exist", logx.LogField{Key: "shortUrl", Value: req.ShortUrl})
		return nil, Err404
	}

	fmt.Println("开始查询redis缓存和DB...")
	// 1.1 查询数据库之前增加缓存层

	// go-zero的缓存支持singleflight
	u, err := l.svcCtx.ShortUrlMapModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortUrl, Valid: true})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			logx.Errorf("shortUrl not found:%s", req.ShortUrl)
			return nil, errors.New("404")
		}
		logx.Errorw("ShortUrlMapModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()},
			logx.LogField{Key: "shortUrl", Value: req.ShortUrl})
		return nil, err
	}

	// 2. 返回查询到的长链接，在handler层返回重定向响应
	return &types.ShowResponse{LongUrl: u.Lurl.String}, nil
}
