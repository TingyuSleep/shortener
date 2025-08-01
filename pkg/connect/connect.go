package connect

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

// client 全局的http客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get 判断url是否能请求通
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect client.Get failed", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
		return false
	}
	resp.Body.Close()
	// 200才算通过
	return resp.StatusCode == http.StatusOK // 别人给我发一个跳转响应（301状态码）这里也不算通过
}
