package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	route := gin.Default()
	// Ping test
	route.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	route.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		log.Println("user ==>:", user)
		value, ok := DB[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// 表单处理
	route.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		name := c.DefaultPostForm("name", "jack")
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": message,
			"name":    name,
		})
	})
	// 测试
	// 1:  curl -d "name=polly&message=helloworld" http://localhost:8080/form_post

	// JSON 处理
	type JSON struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}
	route.POST("/json_post", func(c *gin.Context) {
		var jsonb JSON
		if err := c.ShouldBindJSON(&jsonb); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("name:", jsonb.Name, "message:", jsonb.Message)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	// 测试:
	// 1. curl -H "Content-Type:application/json" -X POST -d '{"name":"polly","message":"roll the world"}' http://localhost:8080/json_post

	route.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0") // 默认为0
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s, page: %s, name: %s, message: %s\n", id, page, name, message)
	})
	// 测试 :
	// curl  -d "name=maliuliu&message=hahahah&page=2" http://localhost:8080/post\?id\=1234

	// 文件上传
	route.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		fmt.Println("filename:", file.Filename, "size:", file.Size, "header:", file.Header)
		c.String(http.StatusOK, "upload OK"+file.Filename)
	})
	return route
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
