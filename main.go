package main

import (
	"code2/api/router"
	_ "code2/common/config"
	"code2/common/database"
	_ "code2/common/database"
	"code2/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

func main() {

	//linkDb()
	runGin()
}

type User struct {
	Name     string `json:"name"`
	Accout   string `json:"accout"`
	Password string `json:"password"`
}

func init() {
	_ = database.DB.AutoMigrate(&model.Account{})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func runGin() {
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	// 加载图标
	app.Use(favicon.New("./assets/favicon.png"))
	// 允许跨域
	app.Use(Cors())

	router.CollectRoute(app)

	app.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	app.GET("/user", func(context *gin.Context) {

		id := context.Query("id")
		name := context.Query("name")
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"id":   id,
				"name": name,
			},
		})

	})

	app.GET("/user/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"id":   id,
				"name": name,
			},
		})
	})
	// 添加用户
	app.POST("/user/create", func(c *gin.Context) {
		// 获取body对象
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
	})

	port := viper.GetString("server.port")
	err := app.Run(":" + port)

	if err != nil {
		println(`服务器启动失败`)
	} else {
		println(`服务器启动成功：http://localhost:` + port)
	}
}

// 连接数据库
func linkDb() {

	var dsn = "root:mmbb1234@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=true"
	var err error
	db, _ = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 设置最大的连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大的连接数
	db.SetMaxOpenConns(10)
	// 空闲的连接数
	db.SetMaxIdleConns(10)

	fmt.Printf("%v", db)

}

// 添加一个用户
func (u User) insertUser() bool {
	s := "insert into user (name,accout,password) values (?,?,?)"
	exec, err := db.Exec(s, u.Name, u.Accout, u.Password)
	if err == nil {
		id, _ := exec.LastInsertId()
		log.Printf("插入成功，id为%v", id)
		return true
	}
	return false
}
