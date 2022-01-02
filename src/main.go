package main

import (
	"log"

	config "github.com/airaider/go_db/config"
	routes "github.com/airaider/go_db/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connect DB
	config.Connect()

	//Init Router
	router := gin.Default()

	//Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
