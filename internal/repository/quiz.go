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

func (qr *QuizRepository) Create(c context.Context, quiz models.Quiz, userID uint) (int, error) {
	var id int
	query := `INSERT INTO quizes(
		userid, title, qcount)
		VALUES ($1, $2, $3) returning id;`
	err := qr.db.QueryRow(c, query, userID, quiz.Title, quiz.CountOfQuestion).Scan(&id)
	if err != nil {
		return 0, err
	}
	for _, question := range quiz.Questions {
		var questionID int
		questionQuery := `INSERT INTO question(
			quizid, question, answer)
			VALUES ($1, $2, $3) returning id;`
		err := qr.db.QueryRow(c, questionQuery, id, question.Title, question.CorrectAnswer).Scan(&questionID)
		if err != nil {
			return 0, err
		}
		for _, variant := range question.Variants {
			variantsQuery := `INSERT INTO variants(
				questionid, variant)
				VALUES ($1, $2);`
			_, err := qr.db.Exec(c, variantsQuery, questionID, variant.Title)
			if err != nil {
				return 0, err
			}
		}

	}
	return id, nil
}

func (qr *QuizRepository) GetAllAdmin(c context.Context, userID uint) ([]models.Quiz, error) {
	quiz := []models.Quiz{}
	return quiz, nil
}

func (qr *QuizRepository) GetAllUser(c context.Context, userID uint) ([]models.Quiz, error) {
	quiz := []models.Quiz{}
	return quiz, nil
}

func (qr *QuizRepository) GetByIDAdmin(c context.Context, quizID int) (models.Quiz, error) {
	quiz := models.Quiz{}
	query := `SELECT id, title, qcount, passed FROM quizes where id = $1;`
	err := qr.db.QueryRow(c, query, quizID).Scan(&quiz.ID, &quiz.Title, &quiz.CountOfQuestion, &quiz.PassedCount)
	if err != nil {
		return quiz, err
	}
	questions := []models.Question{}
	questionQuery := `SELECT id, question, answer FROM question where quizid = $1 order by orderid;`
	rows, err := qr.db.Query(c, questionQuery, quiz.ID)
	if err != nil {
		return quiz, err
	}
	for rows.Next() {
		question := models.Question{}
		err := rows.Scan(&question.ID, &question.Title, &question.CorrectAnswer)
		if err != nil {
			return quiz, err
		}
		variants := []models.Variant{}
		variantsQuery := `SELECT variant FROM variants where questionid = $1 order by orderid;`
		rowss, err := qr.db.Query(c, variantsQuery, question.ID)
		if err != nil {
			return quiz, err
		}
		for rowss.Next() {
			variant := models.Variant{}
			err := rowss.Scan(&variant.Title)
			if err != nil {
				return quiz, err
			}
			variants = append(variants, variant)
		}
		question.Variants = variants
		questions = append(questions, question)
	}
	quiz.Questions = questions
	return quiz, nil
}

func (qr *QuizRepository) GetByIDUser(c context.Context, quizID int  , userID uint) (models.Quiz, error) {
	quiz := models.Quiz{}
	query := `SELECT q.id, q.title, q.qcount, case when r.userID != 0 then true else false end FROM quizes q , results r where q.id = $1 and r.quizid = q.id and r.userid = $2;`
	err := qr.db.QueryRow(c, query, quizID , userID).Scan(&quiz.ID, &quiz.Title, &quiz.CountOfQuestion, &quiz.IsPassed)
	if err != nil {
		return quiz, err
	}
	questions := []models.Question{}
	questionQuery := `SELECT id, question, answer FROM question where quizid = $1 order by orderid;`
	rows, err := qr.db.Query(c, questionQuery, quiz.ID)
	if err != nil {
		return quiz, err
	}
	for rows.Next() {
		question := models.Question{}
		err := rows.Scan(&question.ID, &question.Title, &question.CorrectAnswer)
		if err != nil {
			return quiz, err
		}
		variants := []models.Variant{}
		variantsQuery := `SELECT variant FROM variants where quizid = $1 order by orderid;`
		rowss, err := qr.db.Query(c, variantsQuery, quiz.ID)
		if err != nil {
			return quiz, err
		}
		for rowss.Next() {
			variant := models.Variant{}
			err := rowss.Scan(&variant.Title)
			if err != nil {
				return quiz, err
			}
			variants = append(variants, variant)
		}
		question.Variants = variants
		questions = append(questions, question)
	}
	quiz.Questions = questions
	return quiz, nil
}

func (qr *QuizRepository) Delete(c context.Context, quizID int) error {
	return nil
}
