package main

import (
	"fmt"
	"log"

	"github.com/schoukri/joke-server/service"
	"github.com/spf13/viper"
)

func init() {
	// read config.yaml file
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s\n", err)
	}

	// set defaults
	viper.SetDefault("person_service_base_url", "https://names.mcquay.me/api/v0")
	viper.SetDefault("joke_service_base_url", "https://api.icndb.com/jokes")
	viper.SetDefault("http_server_port", 5000)
}

func main() {

	// initialize app
	app := NewApp(
		service.NewPersonService(viper.GetString("person_service_base_url")),
		//service.NewMockPersonService(),
		service.NewJokeService(viper.GetString("joke_service_base_url")),
	)

	// start app
	addr := fmt.Sprintf(":%s", viper.GetString("http_server_port"))
	log.Printf("Listening on addr %s\n", addr)
	log.Fatal(app.Start(addr))
}
