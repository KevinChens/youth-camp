package service

import (
	"errors"
	"fmt"
	"lesson2/project-demo2/repository"
	"sync"
)

type TopicInfo struct {
	Topic *repository.Topic
	User  *repository.User
}

type PostInfo struct {
	Post *repository.Post
	User *repository.User
}

type PageInfo struct {
	TopicInfo *TopicInfo
	PostList  []*PostInfo
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic   *repository.Topic
	posts   []*repository.Post
	userMap map[int64]*repository.User
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topicId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topicId,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {

	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	errChan := make(chan error, 2)
	go func() {
		defer wg.Done()
		topic, err := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		if err != nil {
			errChan <- err
			return
		}
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts, err := repository.NewPostDaoInstance().QueryPostByParentId(f.topicId)
		if err != nil {
			errChan <- err
			return
		}
		f.posts = posts
	}()
	wg.Wait()

	select {
	case err := <-errChan:
		return err
	default:
	}
	uids := []int64{f.topic.Id}
	for _, post := range f.posts {
		uids = append(uids, post.Id)
	}
	userMap, err := repository.NewUserDaoInstance().MQueryUserById(uids)
	if err != nil {
		return err
	}
	f.userMap = userMap
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	userMap := f.userMap
	topicUser, ok := userMap[f.topic.UserId]
	if !ok {
		errors.New("has no topic user info")
	}
	postList := make([]*PostInfo, 0)
	for _, post := range f.posts {
		postUser, ok := userMap[post.UserId]
		if !ok {
			errors.New("has no post user info for " + fmt.Sprint(post.UserId))
		}
		postList = append(postList, &PostInfo{
			Post: post,
			User: postUser,
		})
	}
	f.pageInfo = &PageInfo{
		TopicInfo: &TopicInfo{
			Topic: f.topic,
			User:  topicUser,
		},
		PostList: postList,
	}
	return nil
}
