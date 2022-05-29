package repository

import "sync"

// Topic information
type Topic struct {
	Id         int64  `json:"id"`
	CreateTime int64  `json:"create_time"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

// TopicDao ?
type TopicDao struct {
}

var (
	topicDao *TopicDao
	// singleton, space saving
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
