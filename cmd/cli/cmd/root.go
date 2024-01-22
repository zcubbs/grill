package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/cmd/config"
	"github.com/zcubbs/grill/cmd/cli/cmd/create"
	"github.com/zcubbs/grill/cmd/cli/cmd/get"
	"github.com/zcubbs/grill/cmd/cli/cmd/login"
	"github.com/zcubbs/grill/cmd/cli/cmd/logout"
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
	rootCmd.DisableSuggestions = true
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(aboutCmd)
	rootCmd.AddCommand(ping.Cmd)
	rootCmd.AddCommand(login.Cmd)
	rootCmd.AddCommand(logout.Cmd)
	rootCmd.AddCommand(create.Cmd)
	rootCmd.AddCommand(get.Cmd)
	rootCmd.AddCommand(config.Cmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "configFilePath",
		fmt.Sprintf("%s%d.%s.yaml", os.Getenv("HOME"), os.PathSeparator, cliName),
		"config file (default is $HOME/.grill.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
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
