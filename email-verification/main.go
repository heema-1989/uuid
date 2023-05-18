package main

import (
	"email-verification/initializers"
	_ "email-verification/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	beego.Run()
}
