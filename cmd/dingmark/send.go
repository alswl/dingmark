package main

import (
	"errors"
	"fmt"
	robot "github.com/iaping/go-dingtalk-robot"
	"github.com/iaping/go-dingtalk-robot/message"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sendCmd = &cobra.Command{
	Use: "send",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("requires 2 argument, title and text")
		}
		token := viper.GetString("token")
		if token == "" {
			return fmt.Errorf("%s is blank", "token")
		}
		secret := viper.GetString("secret")
		if secret == "" {
			return fmt.Errorf("%s is blank", "secret")
		}
		if args[0] == "" || args[1] == "" {
			return fmt.Errorf("invalid args input")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		secret := viper.GetString("secret")
		title := args[0]
		text := args[1]

		//机器人Token是webhook上的access_token参数值
		client := robot.New(token, secret)

		//markdown类型
		markdown := message.NewMarkdown()
		markdown.SetTitle(title)
		markdown.SetText(text)
		resp, err := client.Send(markdown)
		if err == nil {
			fmt.Println("result:", resp.IsSuccess(), "code:", resp.GetCode(), "message:", resp.GetMessage())
		}

	},
}
