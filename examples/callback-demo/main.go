package main

import (
	"flag"
	"fmt"
	goqywechat "github.com/xqk/go-qywechat"
	"net/http"
)

type dummyRxMessageHandler struct{}

// OnIncomingMessage 一条消息到来时的回调。
func (m *dummyRxMessageHandler) OnIncomingMessage(msg *goqywechat.RxMessage) error {
	// You can do much more!
	fmt.Printf("传入消息: %s\n", msg)
	return nil
}

func main() {
	pAddr := flag.String("addr", "[::]:8000", "要侦听的地址和端口")
	pToken := flag.String("token", "", "企业微信配置的Token")
	pEncodingAESKey := flag.String("key", "", "企业微信配置的EncodingAESKey")

	flag.Parse()

	hh, err := goqywechat.NewHTTPHandler(*pToken, *pEncodingAESKey, &dummyRxMessageHandler{})
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", hh)

	err = http.ListenAndServe(*pAddr, mux)
	if err != nil {
		panic(err)
	}
}
