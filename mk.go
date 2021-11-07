package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main2() {
	r := gin.Default()

	r.GET("/ping", ping())

	r.GET("user/:name/*action", test1())

	r.GET("user", user())

	r.POST("form", form)

	// 路由分组
	v1 := r.Group("v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("v2")
	{
		v2.GET("/login", login)
		v2.GET("/submit", submit)
	}

	if err := r.Run(":9000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Gin",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user", helloHandler)
	return r
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func ping() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "success",
		})
	}
}

func test1() func(context *gin.Context) {
	return func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		fmt.Println(action)
		// 截取/
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	}
}

func user() func(context *gin.Context) {
	return func(context *gin.Context) {
		name := context.DefaultQuery("name", "zhangsan")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	}
}

func form(context *gin.Context) {
	types := context.DefaultPostForm("type", "post")
	username := context.PostForm("username")
	password := context.PostForm("password")

	context.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, password, types))
}
