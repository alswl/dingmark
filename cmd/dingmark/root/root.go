package root

import (
	"github.com/alswl/dingmark/pkg/version"
	"github.com/spf13/cobra"
)

const App = "dingmark"

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:              App,
		TraverseChildren: true,
		Version:          version.Version,
	}
	root.AddCommand(SendCmd)
	return root
}
