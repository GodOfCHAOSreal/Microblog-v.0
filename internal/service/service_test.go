package service

import (
	"MicroBlog_v0/internal/models"
	"fmt"
	"testing"
)

func TestRegistration(t *testing.T) {
	userBase := models.NewUserBase()
	postBase := models.NewPostBase()

	blog := NewBlog(userBase, postBase)
	err := blog.Registration("Antony")
	if err != nil {
	}

	if _, ok := blog.userBase.Users["Antony"]; !ok {
		t.Errorf("Тест не пройден, данная запись не зарегана!")
	}
}

func TestAuthorize(t *testing.T) {
	userBase := models.NewUserBase()
	postBase := models.NewPostBase()

	blog := NewBlog(userBase, postBase)
	err := blog.Registration("Antony")
	if err != nil {
		fmt.Println("Ошибка!")
	}

	result, err := blog.Authorize("Antony")
	if err != nil {
	}
	result1, err1 := blog.Authorize("Valera")
	if err1 != nil {
	}
	if result != true || result1 != false {
		t.Errorf("Тест не пройден, функция авторизации не работает!")
	}
}

func TestCreateNewPost(t *testing.T) {
	userBase := &models.UserBase{
		Users: map[string]models.User{
			"Antony": {
				Name: "Antony",
				ID:   "K23In",
			},
			"Misha": {
				Name: "Misha",
				ID:   "ZFm6E",
			},
		},
	}
	postBase := &models.PostBase{
		Posts: []models.Post{
			{
				Author: models.User{
					Name: "Antony",
					ID:   "K23In",
				},
				Text:  "Всем привет!",
				MsgID: "K123dFdsI2",
				Likes: 2,
			},
		},
	}

	blog := NewBlog(userBase, postBase)

	err := blog.CreateNewPost(blog.userBase.Users["Misha"], "Нияз крутой!")
	if err != nil {
	}

	expected := "Нияз крутой!"
	for _, value := range blog.postBase.Posts {
		if value.Text == expected {
			return
		}
	}
	t.Errorf("Тест не пройден, функция создания поста не работает!")

}

func TestGetAllPosts(t *testing.T) {
	userBase := &models.UserBase{
		Users: map[string]models.User{
			"Antony": {
				Name: "Antony",
				ID:   "K23In",
			},
			"Misha": {
				Name: "Misha",
				ID:   "ZFm6E",
			},
		},
	}
	postBase := &models.PostBase{
		Posts: []models.Post{
			{
				Author: models.User{
					Name: "Antony",
					ID:   "K23In",
				},
				Text:  "Всем привет!",
				MsgID: "K123dFdsI2",
				Likes: 2,
			},
			{
				Author: models.User{
					Name: "Misha",
					ID:   "ZFm6E",
				},
				Text:  "Всё круто!",
				MsgID: "J54Fdsa21d",
				Likes: 23,
			},
		},
	}

	blog := NewBlog(userBase, postBase)
	posts := blog.GetAllPosts()

	if posts[0] != "Всем привет!" || posts[1] != "Всё круто!" {
		t.Errorf("Тест не пройден, функция получения постов не работает!")
	}
}

func TestLikeByMsgID(t *testing.T) {
	userBase := &models.UserBase{
		Users: map[string]models.User{
			"Antony": {
				Name: "Antony",
				ID:   "K23In",
			},
			"Misha": {
				Name: "Misha",
				ID:   "ZFm6E",
			},
		},
	}
	postBase := &models.PostBase{
		Posts: []models.Post{
			{
				Author: models.User{
					Name: "Antony",
					ID:   "K23In",
				},
				Text:  "Всем привет!",
				MsgID: "K123dFdsI2",
				Likes: 2,
			},
			{
				Author: models.User{
					Name: "Misha",
					ID:   "ZFm6E",
				},
				Text:  "Всё круто!",
				MsgID: "J54Fdsa21d",
				Likes: 23,
			},
		},
	}

	blog := NewBlog(userBase, postBase)
	if err := blog.LikeByMsgID("J54Fdsa21d"); err != nil {
	}
	expected := 24
	if blog.postBase.Posts[1].Likes != expected {
		t.Errorf("Тест не пройден, функция лайков не работает!")
	}
}
