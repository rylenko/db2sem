package repo

import (
	"db2sem/internal/db/pg"
)

type Repo struct {
	conn pg.Conn
}

func New(conn pg.Conn) *Repo {
	return &Repo{conn: conn}
}
