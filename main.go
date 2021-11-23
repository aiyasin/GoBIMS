package main

import (
	"GoBIMS/model"
	"GoBIMS/routes"
	"GoBIMS/utils"
)

func main() {
	model.InitDB()
	router := routes.InitRouter()
	router.Run(utils.HttpPort)
}
