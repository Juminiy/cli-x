package cmd 

import (
	"fmt" 
	"github.com/spf13/viper"
) 

func initConfig() { 
	viper.SetConfigName("cli-x")
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