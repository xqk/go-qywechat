package go_qywechat

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/url"
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

		app.SpawnAccessTokenRefresher()
	})

	convey.Convey("获取客户详情", t, func() {
		resp, err := app.GetExternalContact("wmWtJeDgAASAeTFXxuvrOOkw0GpqLZgw")
		convey.Convey("", func() {
			convey.So(err, convey.ShouldBeNil)
		})
		fmt.Println(resp)
	})

	convey.Convey("获取客户列表", t, func() {
		resp, err := app.ListExternalContact("xiaqiankun")
		convey.Convey("", func() {
			convey.So(err, convey.ShouldBeNil)
		})
		fmt.Println(resp)
	})
}

func TestQyWechatSystemApp_ExternalContactAddContact(t *testing.T) {
	corpId := "ww38723bfa8894c18f"
	appSecret := "2UleL58-Zk2tYdMVicPj8_XhRjDJlhHDyhDCllLos9k"

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("socks5://192.168.0.17:1082")
	}

	client := New(
		corpId,
		WithCache("47.101.69.114:63791", "9ac313088d7b42005a7ff6329b57e44943cfd3a1", 0),
		WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				Proxy: proxy,
			},
		}),
	)
	app := client.WithSystemApp(appSecret)

	typ := 1
	scene := 2
	style := 1
	remark := "zhukePro"
	state := "zhukePro-dev-xqk"
	user := []string{"sean"}
	party := []int{}
	expiresIn := 0
	chatExpiresIn := 0
	unionID := ""
	var conclusions Conclusions
	contact, err := app.ExternalContactAddContact(typ, scene, style, remark, true, state, user, party, false, expiresIn, chatExpiresIn, unionID, conclusions)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(contact.ConfigID)
	fmt.Println(contact.QRCode)

	// 删除
	err = app.ExternalContactDelContactWay(contact.ConfigID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("删除成功", contact.ConfigID)
}

func TestQyWechatSystemApp_ExternalContactGetContactWay(t *testing.T) {
	corpId := "ww38723bfa8894c18f"
	appSecret := "2UleL58-Zk2tYdMVicPj8_XhRjDJlhHDyhDCllLos9k"

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("socks5://192.168.0.17:1082")
	}

	client := New(
		corpId,
		WithCache("47.101.69.114:63791", "9ac313088d7b42005a7ff6329b57e44943cfd3a1", 0),
		WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				Proxy: proxy,
			},
		}),
	)
	app := client.WithSystemApp(appSecret)

	way, err := app.ExternalContactGetContactWay("2ecbddec3ff74b9785722e0b020b7584")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(way.QRCode)
}

func TestQyWechatSystemApp_ExternalContactDelContactWay(t *testing.T) {
	corpId := "ww38723bfa8894c18f"
	appSecret := "2UleL58-Zk2tYdMVicPj8_XhRjDJlhHDyhDCllLos9k"

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("socks5://192.168.0.17:1082")
	}

	client := New(
		corpId,
		WithCache("47.101.69.114:63791", "9ac313088d7b42005a7ff6329b57e44943cfd3a1", 0),
		WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				Proxy: proxy,
			},
		}),
	)
	app := client.WithSystemApp(appSecret)

	err := app.ExternalContactDelContactWay("2ecbddec3ff74b9785722e0b020b7584")
	if err != nil {
		t.Error(err)
		return
	}
}
