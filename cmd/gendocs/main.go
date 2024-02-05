package main

import (
	"github.com/alswl/dingmark/cmd/dingmark/root"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(root.NewRootCmd(), "./docs/commands")
	if err != nil {
		logrus.WithError(err).Fatal("generate markdown failed")
	}
}
