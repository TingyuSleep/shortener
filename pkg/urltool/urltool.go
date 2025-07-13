package urltool

import (
	"errors"
	"net/url"
	"path"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	if len(myUrl.Host) == 0 {
		return "", errors.New("no host in targetUrl")
	}

	// path.Base() 用于获取路径的最后一个元素
	return path.Base(myUrl.Path), nil
}
