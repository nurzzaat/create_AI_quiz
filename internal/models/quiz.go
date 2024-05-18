package models

import "context"

type Quiz struct {
	ID              int        `json:"id"`
	Title           string     `json:"title"`
	Questions       []Question `json:"questions"`
	CountOfQuestion int        `json:"countOfQuestion"`
	PassedCount     int        `json:"passedCount"`
	IsPassed        bool       `json:"isPassed"`
	Points          int        `json:"points"`
	Speciality      string     `json:"speciality"`
	Timer           string     `json:"timer"`
}

type Question struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Variants      []Variant `json:"variants"`
	CorrectAnswer string    `json:"correctAnswer"`
}

type Variant struct {
	Title string `json:"title"`
}

type Submission struct {
	Points  int      `json:"points"`
	Answers []string `json:"answers"`
	Answer  string   `json:"-"`
	Timer   string   `json:"timer"`
}
type StudentResult struct {
	ID        int        `json:"userId"`
	Email     string     `json:"email"`
	Questions []Question `json:"questions"`
	Point     int        `json:"point"`
	Answers   []string   `json:"answers"`
	Percent   int        `json:"percent"`
	Timer     string     `json:"timer"`
}

type QuizRepository interface {
	Create(c context.Context, quiz Quiz, userID uint) (int, error)
	GetAllAdmin(c context.Context, userID uint, search string) ([]Quiz, error)
	GetAllUser(c context.Context, userID uint, search string) ([]Quiz, error)
	GetByIDAdmin(c context.Context, quizID int) (Quiz, error)
	GetByIDUser(c context.Context, quizID int, userID uint) (Quiz, error)
	Delete(c context.Context, quizID int) error

	Submit(c context.Context, quizID int, userID uint, submission Submission) error
	GetStudentsByQuizID(c context.Context, quizID int) ([]UserQuiz, error)
	GetPermittedStudentsByQuizID(c context.Context, quizID int) ([]User, error)
	AddStudentToQuiz(c context.Context, quizID int, userID int) error
	DeleteStudentFromQuiz(c context.Context, quizID int, userID int) error
	GetStudentResult(c context.Context, quizID int, userID int) (StudentResult, error)
}
