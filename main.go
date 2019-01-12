package main

import (
	"fmt"
	"github.com/sundogrd/content-api/middlewares/sdsession"
	"github.com/sundogrd/content-api/utils/config"
	"os"

	"github.com/sundogrd/content-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/utils/db"
	"github.com/sundogrd/content-api/utils/redis"
)

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	var err error
	config.Init()
	name := config.GetString("name")
	port := config.GetString("port")

	// 初始化默认redis db, 后面在使用的时候import "github.com/ihahoo/go-api-lib/redis" 通过redis.DB(0)调用实例
	// 如果要初始化多个redis的db，则在这里添加，比如redis.Init(1)就建立了一个db 1的连接
	// 如果不使用redis，则删除这里以及其它和redis相关的包引入
	err = redis.Init(0)
	if err != nil {
		fmt.Printf("[Main] Init Redis error: %+v", err)
		os.Exit(1)
	}

	// 初始化数据库
	dbClient, err := db.Init()
	if err != nil {
		fmt.Printf("[Main] Init DB error: %+v", err)
		os.Exit(1)
	}
	defer dbClient.Close()

	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(CORSMiddleware())
	r.Use(sdsession.Middleware(`{
		"storeDriver":"redis",
		"cookieName":"session_id",
		"EnableSetCookie":true,
		"secure":false,
		"cookieLifeTime": 86400,
		"Domain":""
	}`))

	routes.Routes(r)

	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"errcode": 405, "errmsg": "Method Not Allowed"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"errcode": 404, "errmsg": "Not Found"})
	})

	fmt.Println(name + " start listening at http://localhost:" + port)
	fmt.Printf("==> 🚀 %s listening at %s\n", name, port)
	r.Run(":" + port)
}
