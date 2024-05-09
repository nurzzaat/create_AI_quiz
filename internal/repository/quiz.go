package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

type QuizRepository struct {
	db *pgxpool.Pool
}

func NewQuizRepository(db *pgxpool.Pool) models.QuizRepository {
	return &QuizRepository{db: db}
}

func (qr *QuizRepository) Create(c context.Context) {

}
