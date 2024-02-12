//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"github.com/alswl/dingmark/pkg/services"
	"github.com/alswl/dingmark/pkg/tokens"
	robot "github.com/alswl/dingmark/third_party/go-ding-robot"
	"syscall/js"
)

func init() {
	// https://corsproxy.github.io/
	// did not support POST

	// https://cors.sh
	// token is required
	robot.Webhook = "https://proxy.cors.sh/https://oapi.dingtalk.com/robot/send"

	//https://cors-proxy.htmldriven.com/
	//robot.Webhook = "https://cors-proxy.htmldriven.com/?url=https://oapi.dingtalk.com/robot/send"
}

func jsFunc(this js.Value, args []js.Value) interface{} {
	if len(args) != 4 {
		return js.ValueOf("args length must be 4")
	}
	token := args[0]
	secret := args[1]
	title := args[2]
	text := args[3]
	robot.ExtendHeaders = map[string]string{
		"x-cors-api-key": tokens.CORSDotSH,
	}

	resp, err := services.SendMarkdown(
		token.String(),
		secret.String(),
		title.String(),
		text.String(),
	)
	if err != nil {
		return js.ValueOf(err.Error())
	}
	if resp == nil {
		return js.ValueOf(nil)
	}
	marshal, err := json.Marshal(resp)
	return js.ValueOf(string(marshal))
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("SendMarkdown", js.FuncOf(jsFunc))
	<-done
}
