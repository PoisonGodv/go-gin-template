package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test_wxlogin/service"
)

func Router() *gin.Engine {
	fmt.Print("进入路由了")
	r := gin.Default()
	//用户模块
	userGroup := r.Group("/user")
	{
		userGroup.GET("/getUserList", service.GetUserList)
		userGroup.POST("/createUser", service.CreateUser)
		userGroup.DELETE("/deleteUser", service.DeleteUser)
		userGroup.PUT("/updateUser", service.UpdateUser)
	}

	return r
}
