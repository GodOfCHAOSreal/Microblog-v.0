package models

type User struct {
	name string
	ID   string
}

type Post struct {
	author User
	text   string
	msgID  string
	likes  map[string]User
}
