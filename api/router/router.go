package router

import (
	"code2/api/user"
	"github.com/gin-gonic/gin"
)

func CollectRoute(router *gin.Engine) {

	userGroup := router.Group("/user")

	user.RegisterRouter(userGroup)

}
