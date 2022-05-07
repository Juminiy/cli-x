package cmd 

import (
	"fmt"
  
	"github.com/spf13/cobra"
)

var VersionCmd = & cobra.Command {
	Use: "version" ,
	Short: "Cli-x version" ,
	Long: `Cli-x Version Number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chisato Cli Version 1.0") 
	},
}