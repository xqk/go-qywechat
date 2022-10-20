package go_qywechat

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQyWechatNew(t *testing.T) {
	convey.Convey("不带参数构造 QyWechat 实例", t, func() {
		corpId := "ww38723bfa8894c18f"
		client := New(corpId)

		convey.Convey("corpId 应该正确设置了", func() {
			convey.So(client.CorpID, convey.ShouldEqual, "ww38723bfa8894c18f")
		})
	})
}

func TestQyWechat_WithSystemApp(t *testing.T) {
	convey.Convey("客户列表", t, func() {
		corpId := "ww38723bfa8894c18f"
		appSecret := "2UleL58-Zk2tYdMVicPj8_XhRjDJlhHDyhDCllLos9k"
		client := New(corpId, WithCache("192.168.0.17:6382", "", 0))

		app := client.WithSystemApp(appSecret)
		convey.Convey("不能有错", func() {
			accessToken, err := app.getAccessToken()
			convey.So(err, convey.ShouldBeNil)
			convey.So(accessToken, convey.ShouldNotBeNil)
		})
	})
}
