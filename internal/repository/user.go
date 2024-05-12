package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) models.UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(c context.Context, user models.UserRequest) (int, error) {
	var userID int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userQuery := `INSERT INTO users(
		email, password, firstname, lastname, createdat , roleid)
		VALUES ($1, $2, $3, $4, $5 , $6) returning id;`
	err := ur.db.QueryRow(c, userQuery, user.Email, user.Password, user.FirstName, user.LastName, currentTime, user.RoleID).Scan(&userID)
	if err != nil {

		return 0, err
	}
	return userID, nil
}

func (ur *UserRepository) EditUser(c context.Context, user models.User) (int, error) {
	userQuery := `UPDATE users SET email=$1, firstname=$2 , lastname=$3 WHERE id = $4`
	_, err := ur.db.Exec(c, userQuery, user.Email, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil
}
func (ur *UserRepository) DeleteUser(c context.Context, userID int) error {
	query := `delete from users where id = $1`
	_, err := ur.db.Exec(c, query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(c context.Context, email string) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, firstname , lastname , createdat , roleid FROM users where email = $1`
	row := ur.db.QueryRow(c, query, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.RoleID)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByID(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email, password, firstname , lastname , createdat , roleid FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.RoleID)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetProfile(c context.Context, userID int) (models.User, error) {
	user := models.User{}

	query := `SELECT id, email,  firstname , lastname , createdat , roleid FROM users where id = $1`
	row := ur.db.QueryRow(c, query, userID)
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt , &user.RoleID)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetAll(c context.Context) ([]models.User, error) {
	users := []models.User{}

	query := `SELECT id, email,  firstname , lastname , createdat FROM users where roleid = 2`
	rows, err := ur.db.Query(c, query)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) SetUserPassword(c context.Context, password string, userID int) error {
	query := `update users set password = $1 where id = $2`
	_, err := ur.db.Exec(c, query, password, userID)
	if err != nil {
		return err
	}
	return nil
}
