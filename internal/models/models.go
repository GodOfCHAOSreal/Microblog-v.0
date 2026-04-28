package models

type User struct {
	Name string
	ID   string
}

type Post struct {
	Author User
	Text   string
	MsgID  string
	Likes  int
}

type UserBase struct {
	Users map[string]User
}

type PostBase struct {
	Posts []Post
}

func NewUserBase() *UserBase {
	return &UserBase{
		Users: map[string]User{},
	}
}

func NewPostBase() *PostBase {
	return &PostBase{
		Posts: []Post{},
	}
}
