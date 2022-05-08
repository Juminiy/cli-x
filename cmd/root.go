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
	user string 
	body string
)

func init() { 

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&user,"user","u","","username")
	viper.BindPFlag("user",rootCmd.PersistentFlags().Lookup("app.author.name"))

	rootCmd.AddCommand(InfoCmd)
	rootCmd.AddCommand(VersionCmd) 
	rootCmd.AddCommand(ListCmd) 

	rootCmd.AddCommand(HttpCmd)
	HttpCmd.AddCommand(getCmd)
	HttpCmd.AddCommand(postCmd)
	postCmd.Flags().StringVarP(&body,"body","b","","http post body")
	rootCmd.AddCommand(MongoCmd)
}

var rootCmd = &cobra.Command{
	Use:   "cli-x",
	Short: "😀 Cli-x is a cmd tool.",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😋 😍 🥰 You have installed cli-x successfully !")
		fmt.Println("😛 😜 😝 Please input ' cli-x help ' for usage !")
	},
}



func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

