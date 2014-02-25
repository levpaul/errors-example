package main

import (
	"fmt"
)

type User struct {
	Id   int64
	Name string
}

type UserDBService interface {
	GetUserById(id int64) (*User, error)
	InsertUser(user User) error
}

type UserDBServiceImpl struct {
	// Local vars...
}

func (s UserDBServiceImpl) GetUserById(id int64) (*User, error) {
	return nil, nil
}

func (s UserDBServiceImpl) InsertUser(user User) error {
	return nil
}

func main() {
	var udbs UserDBService = UserDBServiceImpl{}
	fmt.Println(udbs.GetUserById(0))
}
