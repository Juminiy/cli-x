package cmd 

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "😁 Show app information. ",
	Long: UsageLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😉 😙 🤩 😎 😆 The Cli-x application information lists : ") 
		fmt.Println("😉 App's Name: ",viper.Get("app.name"))
		fmt.Println("😙 App's Version: ",viper.GetFloat64("app.version"))
		fmt.Println("🤩 App's Author: ",viper.Get("app.author")) 
		fmt.Println("😎 App's git: ",viper.Get("app.gitAddress"))
		fmt.Println("😆 App's created: ",viper.Get("app.createdTime"))
	},
}