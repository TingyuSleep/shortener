package connect

import (
	c "github.com/smartystreets/goconvey/convey" // 别名导入
	"testing"
)

func TestGet(t *testing.T) {
	c.Convey("基础用例", t, func() {
		url := "https://www.liwenzhou.com/posts/Go/unit-test-5/"
		got := Get(url)

		// 断言
		c.So(got, c.ShouldEqual, true)
		//c.ShouldBeTrue(got)
	})

	c.Convey("请求不通的示例", t, func() {
		url := "posts/Go/unit-test-5/"
		got := Get(url)

		// 断言
		c.So(got, c.ShouldEqual, false)
		//c.ShouldBeTrue(got)
	})
}
