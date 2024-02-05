package root

import (
	"fmt"

	"github.com/alswl/dingmark/pkg/version"
	"github.com/spf13/cobra"
)

const App = "dingmark"

func NewRootCmd(cfgFile *string) *cobra.Command {
	root := &cobra.Command{
		Use:              App,
		TraverseChildren: true,
		Version:          version.Version,
	}
	root.PersistentFlags().StringVar(cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.config/%s.yaml)", App))
	root.PersistentFlags().String("token", "", "token")
	root.PersistentFlags().String("secret", "", "secret")

	root.AddCommand(SendCmd)
	return root
}
