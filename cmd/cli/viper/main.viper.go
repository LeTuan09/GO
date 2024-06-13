package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // name file
	viper.SetConfigType("yaml")

	// * Read file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	// read server configuration
	fmt.Println("Server Port: ", viper.GetInt("server.port"))
	fmt.Println("Server Port: ", viper.GetString("security.jwt.key"))
}
