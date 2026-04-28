package service

import (
	"MicroBlog_v0/internal/models"
)

type Blog struct {
	userBase *models.UserBase
	postBase *models.PostBase
}

func NewBlog(userBase *models.UserBase, postBase *models.PostBase) *Blog {
	return &Blog{
		userBase: userBase,
		postBase: postBase,
	}
}
func (b *Blog) Registration(name string) error {
	if _, ok := b.userBase.Users[name]; ok {
		return ErrorUserExists
	}
	for {
		newID, err := generateID(5)
		if err != nil {
			return err
		}
		idExists := false

		for _, value := range b.userBase.Users {
			if newID == value.ID {
				idExists = true
				break
			}
		}
		if idExists {
			continue
		}

		b.userBase.Users[name] = models.User{Name: name, ID: newID}
		return nil
	}
}

func (b *Blog) Authorize(name string) (bool, error) {
	if _, ok := b.userBase.Users[name]; !ok {
		return false, ErrorUserNotExist
	}

	return true, nil
}

func (b *Blog) CreateNewPost(author models.User, text string) error {
	if _, ok := b.userBase.Users[author.Name]; !ok {
		return ErrorUserNotExist
	}
	for {
		newID, err := generateID(10)
		if err != nil {
			return err
		}
		idExists := false

		for _, value := range b.postBase.Posts {
			if newID == value.MsgID {
				idExists = true
				break
			}
		}
		if idExists {
			continue
		}

		b.postBase.Posts = append(b.postBase.Posts, models.Post{
			Author: author,
			Text:   text,
			MsgID:  newID,
			Likes:  0,
		})
		return nil
	}
}

func (b *Blog) GetAllPosts() []string {
	result := []string{}
	for _, value := range b.postBase.Posts {
		result = append(result, value.Text)
	}
	return result
}

func (b *Blog) LikeByMsgID(msgID string) error {
	for i, value := range b.postBase.Posts {
		if value.MsgID == msgID {
			b.postBase.Posts[i].Likes++
			return nil
		}
	}
	return ErrorPostNotFound
}
