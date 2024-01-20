package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zcubbs/grill/cmd/cli/cmd/ping"
	"os"
)

const cliName = "grill"

var (
	Version string
	Commit  string
	Date    string

	cfgFile string
)

var (
	rootCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getVersion())
		},
	}

	aboutCmd = &cobra.Command{
		Use:   "about",
		Short: "Print the info about this CLI",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			About()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	initConfig()

	rootCmd.DisableSuggestions = true
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(aboutCmd)
	rootCmd.AddCommand(ping.Cmd)

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP("server", "s", "localhost:9000", "The server address in the format of host:port")
	_ = viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
}

func About() {
	figure.NewColorFigure("GRILL", "colossal", "yellow", true).Print()
	figure.NewColorFigure("grill", "morse", "red", true).Print()
	fmt.Println(getFullVersion())
	fmt.Println(getDescription())
	fmt.Println("Copyright (c) 2023 zakaria.elbouwab (zcubbs)")
	fmt.Println("Repository: https://github.com/zcubbs/grill")
}

func getVersion() string {
	return fmt.Sprintf("v%s", Version)
}

func getFullVersion() string {
	return fmt.Sprintf(`
Version: v%s
Commit: %s
Date: %s
`, Version, Commit, Date)
}

func getDescription() string {
	return `
/grill/ is a CLI tool for managing your clusters.
`
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".<config_file_name>" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%s", cliName))
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
