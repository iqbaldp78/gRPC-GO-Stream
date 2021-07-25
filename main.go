package main

import (
	"fmt"
	"grpc-go/modules"
	"grpc-go/modules/services"
	"log"

	"github.com/spf13/viper"
)

func main() {
	setConfig("config")

	log.Printf("Listening on port : %v...\n", viper.GetInt("SERVER_PORT"))
	log.Fatal(modules.ListenGRPC(services.New(), viper.GetInt("SERVER_PORT")))
}

func setConfig(config string) {
	viper.SetConfigName(config)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Configuration file not found: %s", err))
	}

	viper.SetEnvPrefix("HTT")
	viper.AutomaticEnv()
}
