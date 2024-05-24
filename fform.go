package fform

import (
	"database/sql"
	"fform/dialect"
	"fform/log"
	"fform/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not found", driver)
		return nil, nil
	}
	e := &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return e, nil
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.NewSession(e.db, e.dialect)
}
