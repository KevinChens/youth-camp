package controller

import (
	"lesson2/project-demo1/service"
	"strconv"
)

func PublishPost(topicIdStr, content string) *PageData {
	// param convert
	topicId, _ := strconv.ParseInt(topicIdStr, 10, 64)
	postId, err := service.PublishPost(topicId, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"post_id": postId,
		},
	}
}
