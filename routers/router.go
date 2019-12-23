package routers

import (
	"github.com/astaxie/beego"
	"myBeego/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/testPost", &controllers.MainController{})
	//beego.Router("/user", &controllers.UserController{})
	//beego.Router("/getUser", &controllers.UserController{}, "get:GetUser")
	//beego.Router("/hello", &controllers.HelloController{})

	beego.Router("/", &controllers.BlogController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exitLogin", &controllers.ExitLoginController{})

	beego.Router("/category", &controllers.CategoryController{})

}
