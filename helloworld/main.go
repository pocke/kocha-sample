package main

import (
	"github.com/naoina/kocha-sample/helloworld/app/controller"
	"github.com/naoina/kocha-sample/helloworld/config"

	"github.com/naoina/kocha"
)

func main() {
	config.AppConfig.RouteTable = kocha.RouteTable{
		{
			Name:       "root",
			Path:       "/",
			Controller: &controller.Root{},
		},
		{
			Name:       "static",
			Path:       "/*path",
			Controller: &kocha.StaticServe{},
		},
	}
	if err := kocha.Run(config.AppConfig); err != nil {
		panic(err)
	}
}
