package root

import (
	"errors"
	"fmt"

	"github.com/alswl/dingmark/pkg/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SendCmd = &cobra.Command{
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

		resp, err := services.SendMarkdown(token, secret, title, text)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("result:", resp.IsSuccess(), "code:", resp.GetCode(), "message:", resp.GetMessage())
	},
}
