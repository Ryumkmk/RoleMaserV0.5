package main

import (
	"RMV0.5/app/controllers"
	"RMV0.5/app/models"
)

func main() {
	models.GetPjs("1")
	controllers.StartMainServer()
}
