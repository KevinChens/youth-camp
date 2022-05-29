package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lesson2/project-demo1/controller"
	"lesson2/project-demo1/repository"
	"net/http"
)

func main() {
	// 1. init Index
	if err := repository.Init("./data/"); err != nil {
		fmt.Printf("init Index failed, err:%v\n", err)
		return
	}
	// 2. init engine
	r := gin.Default()
	// 3. register router
	r.GET("/community/page/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(http.StatusOK, data)
	})
	r.POST("/community/do", func(c *gin.Context) {
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := controller.PublishPost(topicId, content)
		c.JSON(http.StatusOK, data)
	})
	// 4. run
	err := r.Run(":9091")
	if err != nil {
		fmt.Printf("run failed, err:%v\n", err)
		return
	}
}
