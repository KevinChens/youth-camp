package service

import (
	"errors"
	"lesson2/project-demo1/pkg/snowflake"
	"lesson2/project-demo1/repository"
	"time"
	"unicode/utf16"
)

type PublishPostFlow struct {
	postId  int64
	topicId int64
	content string
}

func PublishPost(topicId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, content).Do()
}

func NewPublishPostFlow(topicId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		topicId: topicId,
		content: content,
	}
}

func (f *PublishPostFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f *PublishPostFlow) checkParam() error {
	if len(utf16.Encode([]rune(f.content))) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	if err := snowflake.Init("2022-05-29", 1); err != nil {
		return err
	}
	id := snowflake.GenID()
	post := &repository.Post{
		Id:         id,
		ParentId:   f.topicId,
		Content:    f.content,
		CreateTime: time.Now().Unix(),
	}
	if err := repository.NewPostDaoInstance().InsertPost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
