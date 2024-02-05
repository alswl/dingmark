package main

import (
	"fmt"
	"os"
	"path"

	"github.com/alswl/dingmark/cmd/dingmark/root"
	"github.com/alswl/dingmark/pkg/version"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
)

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config"))
		viper.SetConfigName(root.App)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	cobra.OnInitialize(initConfig)
	rootCmd := root.NewRootCmd(&cfgFile)

	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	_ = viper.BindPFlag("secret", rootCmd.PersistentFlags().Lookup("secret"))
	// viper args can NOT set required
	rootCmd.SetVersionTemplate(fmt.Sprintf("version {{.Version}}, commit %s, package %s\n", version.Commit, version.Package))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
