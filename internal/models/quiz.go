package models

import "context"

type Quiz struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Questions []Question `json:"questions"`
	CountOfQuestion int `json:"countOfQuestion"`
	PassedCount int `json:"passedCount"`
	IsPassed bool `json:"isPassed,omitempty"`
}

type Question struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Variants []Variant `json:"variants"`
	CorrectAnswer string `json:"correctAnswer"`
}

type Variant struct{
	Title string `json:"title"`
}

type QuizRepository interface {
	Create(c context.Context , quiz Quiz, userID uint)(int , error)
	GetAllAdmin(c context.Context , userID uint)([]Quiz , error)
	GetAllUser(c context.Context , userID uint)([]Quiz , error)
	GetByIDAdmin(c context.Context , quizID int)(Quiz , error)
	GetByIDUser(c context.Context , quizID int , userID uint)(Quiz , error)
	Delete(c context.Context , quizID int)(error)
}