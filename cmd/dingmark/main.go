package main

import (
	"fmt"
	"github.com/alswl/dingmark/pkg/version"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const app = "dingmark"

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
		viper.SetConfigName(app)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

var rootCmd = &cobra.Command{
	Use:              app,
	TraverseChildren: true,
	Version:          version.Version,
}

func main() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.config/%s.yaml)", app))
	rootCmd.PersistentFlags().String("token", "", "token")
	rootCmd.PersistentFlags().String("secret", "", "secret")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("secret", rootCmd.PersistentFlags().Lookup("secret"))
	// viper args can NOT set required
	rootCmd.SetVersionTemplate(fmt.Sprintf("version {{.Version}}, commit %s, package %s\n", version.Commit, version.Package))

	rootCmd.AddCommand(sendCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
