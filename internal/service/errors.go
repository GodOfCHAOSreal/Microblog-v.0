package service

import "errors"

var (
	ErrorNotCorrect   = errors.New("Something wrong!")
	ErrorUserExists   = errors.New("User already exists!")
	ErrorUserNotExist = errors.New("This User does not exist. Registrate first!")
	ErrorPostNotFound = errors.New("This post does not exist! Try again!")
)
