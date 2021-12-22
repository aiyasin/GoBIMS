package main

import (
	"GoBIMS/model"
	"GoBIMS/routes"
	"GoBIMS/utils"
)

func main() {
	model.InitDB()
	router := routes.InitRouter()
	err := router.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
