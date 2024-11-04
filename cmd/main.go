package main

import (
	"fmt"
	"log"

	"github.com/phn00dev/go-URL-Shortener/internal/app"
	"github.com/phn00dev/go-URL-Shortener/internal/setup/constructor"
)

func main() {
	appDependencies, err := app.GetDependencies()
	if err != nil {
		fmt.Println(err)
		return
	}
	constructor.Build(appDependencies)
	appRouter := app.NewApp(appDependencies.Config)
	runServer := fmt.Sprintf("%s:%s", appDependencies.Config.HttpConfig.HttpHost, appDependencies.Config.HttpConfig.HttpPort)
	log.Println(runServer)
	if err := appRouter.Run(runServer); err != nil {
		fmt.Println(err)
		return
	}
}
