package app

import (
	"database/sql"
	"test-app/internal/app/users"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Run() error
}

type server struct {
	rtr *echo.Echo
	cfg *Config
	db  *sql.DB
}

func NewServer(cfg *Config, db *sql.DB) (server, func()) {
	srv := server{
		rtr: echo.New(),
		db:  db,
		cfg: cfg,
	}

	for _, handler := range []Handler{
		NewHandler(srv.db),
		users.NewHandler(users.NewRepository(srv.db)),
	} {
		handler.RegisterRoutes(srv.rtr)
	}

	return srv, func() { srv.rtr.Close() }
}

func (s server) Run() error {
	return s.rtr.Start(":" + s.cfg.Port)
}
