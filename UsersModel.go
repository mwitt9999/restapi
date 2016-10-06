package main

import "time"

type User struct {
	UserId    int
	Username  string
	Name      string
	Password  string
	CreatedAt time.Time
}

type Users []User
