package go_qywechat

import (
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestDefaultOptions(t *testing.T) {
	convey.Convey("给定一个默认值 options", t, func() {
		opts := defaultOptions()
		convey.Convey("opts.HTTP 应该是空值 http.Client", func() {
			convey.So(opts.HTTP, convey.ShouldNotBeNil)
		})

		convey.Convey("opts.QYAPIHost 应该是企业微信官方 API host", func() {
			convey.So(opts.QYAPIHost, convey.ShouldEqual, "https://qyapi.weixin.qq.com")
		})
	})
}

func TestWithHTTPClient(t *testing.T) {
	convey.Convey("给定一个 options", t, func() {
		opts := options{}

		convey.Convey("用 WithHTTPClient 修饰它", func() {
			newClient := http.Client{}
			o := WithHTTPClient(&newClient)
			o.applyTo(&opts)

			convey.Convey("options.HTTP 应该变了", func() {
				convey.So(opts.HTTP, convey.ShouldEqual, &newClient)
			})
		})
	})
}

func TestWithQYAPIHost(t *testing.T) {
	convey.Convey("给定一个 options", t, func() {
		opts := options{}

		convey.Convey("用 WithQYAPIHost 修饰它", func() {
			o := WithQYAPIHost("http://localhost:8000")
			o.applyTo(&opts)

			convey.Convey("options.QYAPIHost 应该变了", func() {
				convey.So(opts.QYAPIHost, convey.ShouldEqual, "http://localhost:8000")
			})
		})
	})
}
