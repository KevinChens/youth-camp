package main

import (
	"github.com/gin-gonic/gin"
	"lesson2/project-demo2/controller"
	"lesson2/project-demo2/pkg/log"
	"lesson2/project-demo2/repository"
	"net/http"
)

func main() {
	// 1. init
	if err := repository.Init(); err != nil {
		return
	}
	if err := log.InitLogger(); err != nil {
		return
	}
	// 2. engine
	r := gin.Default()
	// 3. register router
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/community/page/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(http.StatusOK, data)
	})

	r.POST("/community/do", func(c *gin.Context) {
		uid, _ := c.GetPostForm("uid")
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := controller.PublishPost(uid, topicId, content)
		c.JSON(http.StatusOK, data)
	})
	// 4. run
	err := r.Run(":9091")
	if err != nil {
		return
	}
}
