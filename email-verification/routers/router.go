package routers

import (
	"email-verification/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{}, "get:RegisterUser")
	beego.Router("/welcome", &controllers.RegisterController{}, "post:RegisteredUser")
	beego.Router("/verified", &controllers.RegisterController{}, "get:VerifyRegister")
}
