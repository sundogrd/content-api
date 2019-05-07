package main

import (
	"fmt"
	"github.com/sundogrd/content-api/env"
	"github.com/sundogrd/content-api/middlewares/cors"
	comment2 "github.com/sundogrd/content-api/providers/grpc/comment"
	"os"

	"github.com/sundogrd/content-api/middlewares/sdsession"
	"github.com/sundogrd/content-api/utils/config"
	"github.com/sundogrd/content-api/utils/redis"

	"github.com/sundogrd/content-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/utils/db"
)

func Init() {

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

	commentClient, _, err := comment2.NewGrpcCommentClient()
	if err != nil {
		fmt.Printf("[Main] Init commentClient error: %+v", err)
		os.Exit(1)
	}

	container := env.Container{
		CommentGrpcClient: commentClient,
	}

	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(cors.CORSMiddleware())
	r.Use(sdsession.Middleware(`{
		"storeDriver":"redis",
		"cookieName":"session_id",
		"EnableSetCookie":true,
		"secure":false,
		"cookieLifeTime": 86400,
		"Domain":""
	}`))

	routes.Routes(r, container)

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
