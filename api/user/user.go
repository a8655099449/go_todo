package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加一个用户
func (u User) insertUser() bool {

	return true
}

type User struct {
	Name     string `json:"name"`
	Accout   string `json:"accout"`
	Password string `json:"password"`
}

func RegisterRouter(router *gin.RouterGroup) {
	router.POST("/create", CreateUser)
}

// 创建用户
func CreateUser(c *gin.Context) {
	data, _ := c.GetRawData()
	var user User
	_ = json.Unmarshal(data, &user)

	res := user.insertUser()
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": "ok",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1000,
			"data": "添加失败",
		})
	}
}
