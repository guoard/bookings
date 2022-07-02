package dbrepo

import (
	"database/sql"

	"github.com/guoard/bookings/internal/config"
	"github.com/guoard/bookings/internal/repository"
)

type postgreDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgreRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgreDBRepo{
		App: a,
		DB: conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}

