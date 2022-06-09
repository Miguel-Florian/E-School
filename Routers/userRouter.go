package routers

import (
	controllers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//router.LoadHTMLGlob("Templates/*")
	//All routes related to users comes here
	router.POST("/user/register", controllers.CreateUser())
	router.POST("/user/login", controllers.LoginUser)
	router.GET("/user/:id", controllers.GetUserById())
	router.PUT("/user/:id", controllers.UpdateUserById())
	router.GET("/users", controllers.GetAllUsers())
	router.DELETE("/user/:id", controllers.DeleteUser())
}
