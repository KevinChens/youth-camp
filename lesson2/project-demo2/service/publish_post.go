package service

import (
	"errors"
	"lesson2/project-demo2/pkg/snowflake"
	"lesson2/project-demo2/repository"
	"time"
	"unicode/utf8"
)

type PublishPostFLow struct {
	userId  int64
	content string
	topicId int64
	postId  int64
}

func PublishPost(topicId, userId int64, content string) (int64, error) {
	return NewPublishPostFLow(topicId, userId, content).Do()
}
func NewPublishPostFLow(topicId, userId int64, content string) *PublishPostFLow {
	return &PublishPostFLow{
		userId:  userId,
		content: content,
		topicId: topicId,
	}
}

func (f *PublishPostFLow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f *PublishPostFLow) checkParam() error {
	if f.userId <= 0 {
		return errors.New("userId must be larger than 0")
	}
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFLow) publish() error {
	if err := snowflake.Init("2022-05-29", 1); err != nil {
		return err
	}
	id := snowflake.GenID()
	post := &repository.Post{
		ParentId:   f.topicId,
		UserId:     f.userId,
		Content:    f.content,
		CreateTime: time.Now(),
		Id:         id,
	}
	if err := repository.NewPostDaoInstance().CreatePost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
