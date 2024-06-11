package userhandler

import "server/sql/database"

type User struct {
	Email string `json:"email"`
	Username string `json:"username"`
}


func dbuserToUser (dbUser database.User) User {
	return User{
		Email: dbUser.Email,
		Username: dbUser.Username,
	}
}