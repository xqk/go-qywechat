package go_qywechat

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQyWechatSystemApp_BatchListExternalContact(t *testing.T) {
	corpId := "ww38723bfa8894c18f"
	appSecret := "2UleL58-Zk2tYdMVicPj8_XhRjDJlhHDyhDCllLos9k"
	client := New(corpId, WithCache("192.168.0.17:6382", "", 0))

	app := client.WithSystemApp(appSecret)

	convey.Convey("批量获取客户详情", t, func() {
		resp, err := app.BatchListExternalContact([]string{"xiaqiankun"}, "", 100)

		convey.Convey("", func() {
			convey.So(err, convey.ShouldBeNil)
		})
		for _, v := range resp.Result {
			fmt.Println(v.ExternalContact)
		}
	})
}
