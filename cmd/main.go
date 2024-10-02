package main

import (
	app_config "eslam3400/go-backend-api/config"
	"fmt"
)

func main() {
	fmt.Println("App Starting")
	app_config.LoadEnv()
	app_config.ConnectToDB()
	fmt.Println("App Started")
}
