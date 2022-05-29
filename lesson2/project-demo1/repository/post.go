package repository

import (
	"encoding/json"
	"os"
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	CreateTime int64  `json:"create_time"`
	Content    string `json:"content"`
}

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostByParentIdd(parentId int64) []*Post {
	return postIndexMap[parentId]
}

func (*PostDao) InsertPost(post *Post) error {
	f, err := os.OpenFile("./data/post", os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	marshal, _ := json.Marshal(post)
	if _, err = f.WriteString(string(marshal) + "\n"); err != nil {
		return err
	}

	rwMutex.Lock()
	postList, ok := postIndexMap[post.ParentId]
	if !ok {
		postIndexMap[post.ParentId] = []*Post{post}
	} else {
		postList = append(postList, post)
		postIndexMap[post.ParentId] = postList
	}
	rwMutex.Unlock()
	return nil
}
