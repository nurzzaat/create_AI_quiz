package pkg

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Application struct {
	Env   *Env
	Pql   *pgxpool.Pool
}

func App() (Application, error) {
	app := &Application{}
	app.Env = NewEnv()
	conn, err := NewPgxConn(app.Env)
	if err != nil {
		return Application{}, err
	}
	app.Pql = conn

	return *app, nil
}

func (app *Application) CloseDBConnection() {
	app.CloseDBConnection()
}
