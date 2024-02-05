package main

import (
	"github.com/alswl/dingmark/cmd/dingmark/root"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
)

func main() {
	cfgFile := ""
	err := doc.GenMarkdownTree(root.NewRootCmd(&cfgFile), "./docs")
	if err != nil {
		logrus.WithError(err).Fatal("generate markdown failed")
	}
}
