/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
var (
	UsageLongDescription = "Maybe, we don't need the long description."
)
var rootCmd = &cobra.Command{
	Use:   "cli-x",
	Short: "😀 Cli-x is a cmd tool.",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😋 😍 🥰 You have installed cli-x successfully !")
		fmt.Println("")
	},
}
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "😁 Show app information. ",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😉😙🤩😎😆 The Cli-x application information lists : ") 
		fmt.Println("😉App's Name: ",viper.Get("app.name"))
		fmt.Println("😙App's Version: ",viper.Get("app.version"))
		fmt.Println("🤩App's Author: ",viper.Get("app.author")) 
		fmt.Println("😎App's git: ",viper.Get("app.gitAddress"))
		fmt.Println("😆App's created: ",viper.Get("app.createdTime"))
	},
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() { 
	viper.SetConfigName("cli-x-config")
	viper.SetConfigType("yaml") 
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig() ; err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError) ; ok {
			fmt.Println("Config file not found!")
		} else {
			fmt.Println("Error is ",err)
		}
	} 
}
func init() { 
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(VersionCmd) 
	rootCmd.AddCommand(ServeCmd)
	rootCmd.AddCommand(ListCmd) 
	rootCmd.AddCommand(HttpCmd)
}


