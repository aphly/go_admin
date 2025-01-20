package main

import (
	"go_admin/app"
	"go_admin/app/http/common/router"
)

func main() {
	app.Init()

	//go im.NewHub.Run()
	server := router.Reg()
	err := server.Run(app.Config.Http.Listen)
	if err != nil {
		panic(err)
	}
}
