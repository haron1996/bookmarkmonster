package api

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type BaseHandler struct {
	pool *pgxpool.Pool
}

func NewBaseHandler(pool *pgxpool.Pool) *BaseHandler {
	return &BaseHandler{
		pool: pool,
	}
}
