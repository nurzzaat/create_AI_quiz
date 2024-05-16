package models

import "context"

type User struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	RoleID    uint   `json:"roleId,omitempty"`
	CreatedAt string `json:"createdAt"`
}

type UserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    uint   `json:"roleId"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserQuiz struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Point     int    `json:"point"`
	Percent   int    `json:"percent"`
	Timer string `json:"timer"`
}

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByID(c context.Context, userID int) (User, error)
	GetProfile(c context.Context, userID int) (User, error)
	GetAll(c context.Context) ([]User, error)

	CreateUser(c context.Context, user UserRequest) (int, error)
	EditUser(c context.Context, user User) (int, error)
	DeleteUser(c context.Context, userID int) error
	SetUserPassword(c context.Context, password string, userID int) error
}