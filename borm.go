package Borm

import (
	"database/sql"
	"github.com/odysa/Borm/log"
	"github.com/odysa/Borm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database succeed !")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
		return
	}
	log.Info("Close database succeed!")
}

func (e *Engine) NewSession() *session.Session {
	return session.NewSession(e.db)
}
