package user

import "time"

type User struct {
	Id              string
	Token           string
	CreatedAt       time.Time
	LastInteraction time.Time
	Deleted         bool
}

func GetUserByToken(token string) (*User, error) {}
