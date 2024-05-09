package models

import "context"

type QuizRepository interface {
	Create(c context.Context)

}