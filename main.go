package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/micro-websocket/config"
	"github.com/goboilerplates/micro-websocket/env"
)

func main() {
	env.LoadVariables()
	if env.EnableProdMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	config.SetMiddleWares(router)
	config.SetAPIPaths(router)

	router.Run(env.ServerHost)
}
