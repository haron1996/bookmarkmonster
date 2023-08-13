package api

import (
	"github.com/jackc/pgx/v5"
)

type BaseHandler struct {
	db *pgx.Conn
}

func NewBaseHandler(db *pgx.Conn) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}
