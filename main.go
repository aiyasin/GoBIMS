package main

import (
	"GoBIMS/routes"
)

func main() {
	router := routes.InitRouter()
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
