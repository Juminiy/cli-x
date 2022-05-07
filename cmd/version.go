/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var VersionCmd = & cobra.Command {
	Use: "version" ,
	Short: "😋 Show cli-x version" ,
	Long: `Cli-x Version Number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cli-x Version is ",viper.GetFloat64("app.version")) 
	},
}