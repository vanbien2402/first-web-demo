package main

import (
	"log"

	"github.com/jinzhu/configor"

	"github.com/vanbien2402/first-web-demo/internal/api"
	"github.com/vanbien2402/first-web-demo/internal/api/configs"
)

func main() {
	var config configs.Config
	if err := configor.New(&configor.Config{}).Load(&config); err != nil {
		log.Println("load environment variables failed", err)
		panic(err)
	}
	svr, err := api.NewServer(&config)
	if err != nil {
		log.Println("init server failed", err)
		panic(err)
	}
	if err = svr.Start(); err != nil {
		log.Println("start server failed", err)
		panic(err)
	}
}
