package services

import (
	robot "github.com/alswl/dingmark/third_party/go-ding-robot"
	"github.com/alswl/dingmark/third_party/go-ding-robot/message"
	"github.com/alswl/dingmark/third_party/go-ding-robot/response"
)

func SendMarkdown(token, secret, title, text string) (*response.Response, error) {
	//机器人Token是webhook上的access_token参数值
	client := robot.New(token, secret)

	//markdown类型
	markdown := message.NewMarkdown()
	markdown.SetTitle(title)
	markdown.SetText(text)
	resp, err := client.Send(markdown)
	if err == nil {
		return resp, err
	}
	return resp, nil

}
