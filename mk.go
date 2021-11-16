package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

func testDB() {

	var (
		id   int
		name string
	)

	db, err := sql.Open("mysql",
		"root:12345678@tcp(127.0.0.1:3306)/cas")
	if err != nil {
		log.Fatal(err)
	}
	/**
	我们使用db.Query()将查询发送到数据库。我们像往常一样检查错误。
	我们用defer内置函数推迟了rows.Close()的执行。这个非常重要。
	我们用rows.Next()遍历了数据行。
	我们用rows.Scan()读取每行中的列变量。
	我们完成遍历行之后检查错误。
	*/
	err = db.Ping()
	rows, err := db.Query("select id, gs_name from goods where id = ?", 101)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	if err != nil {
		// do something here
	}

	defer db.Close()
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
