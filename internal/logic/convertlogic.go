package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/pkg/connect"
	"shortener/pkg/md5"
	"shortener/pkg/urltool"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Convert 转链：输入一个长链接 ---> 短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. 校验输入的数据
	// 1.1 数据不能为空，使用validator校验库

	// 1.2 输入的长链接能ping通，是有效的网址
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效的链接")
	}

	// 1.3 判断数据库是否已经转链过（数据库中是否已经存在该长链）
	// 1.3.1 给长链接生成md5值
	md5Value := md5.Sum([]byte(req.LongUrl)) // []byte(字符串) 表示强制类型转换，字符转 -> 字节型切片
	// 1.3.2 拿md5去数据库中查是否存在
	u, err := l.svcCtx.ShortUrlMapModel.FindOneByMd5(l.ctx, sql.NullString{
		String: "md5Value",
		Valid:  true,
	})

	if err != sqlx.ErrNotFound { // 错误不是未查询到记录：即可能查询到记录，也可能是普通的err
		if err == nil { // 说明查到了记录
			return nil, fmt.Errorf("该链接已被转为:%s\n", u.Surl.String)
		}
		// 说明调用时出错了，记录错误日志
		logx.Error("ShortUrlMapModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	// 1.4 输入的不能是一个(数据库中已存在的)短链接（避免循环转链）
	// 输入的是一个完整的的url:q1mi.cn/1d12a
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Error("urltool.GetBasePath failed", logx.LogField{Key: "lurl", Value: req.LongUrl})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlMapModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if err != sqlx.ErrNotFound { // 错误不是未查询到记录：即可能查询到记录，也可能是普通的err
		if err == nil { // 说明查到了记录
			return nil, errors.New("该链接已经是短链了")
		}
		// 说明调用时出错了，记录错误日志
		logx.Error("ShortUrlMapModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	// 2. 取号

	// 3. 号码转链

	// 4. 存储长短链接映射关系

	// 5. 返回响应

	return
}
