/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ck-ghost/pkg"
	"ck-ghost/pkg/common"
	"ck-ghost/pkg/config"
	"ck-ghost/pkg/db"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ck-ghost",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		err := db.InitDB(common.BuildDsn(config.Host, config.Port, config.User, config.Password, config.DBOption))
		if err != nil {
			log.Printf("error: init db failesd, %s\n", err)
			os.Exit(1)
		}
		config.InitConfig()
		err = pkg.Run(context.Background())
		if err != nil {
			log.Printf("error: run func failed, %s\n", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ck-ghost.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&config.Host, "host", "H", "127.0.0.1", "clickhouse server ip")
	rootCmd.Flags().StringVarP(&config.Port, "port", "p", "9000", "clickhouse server ip")
	rootCmd.Flags().StringVarP(&config.User, "username", "u", "ckuser", "clickhouse identity username")
	rootCmd.Flags().StringVarP(&config.Password, "password", "P", "123456", "identity password")
	rootCmd.Flags().StringVarP(&config.DBOption, "dboption", "o", "compress=1&use_client_time_zone=true",
		"clickhouse dsn option")
	rootCmd.Flags().StringSliceVarP(&config.Appids, "appid", "a", []string{}, "project that need to add materialized view")

}
