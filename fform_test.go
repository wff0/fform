package fform

import (
	"errors"
	"fform/session"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/fform"

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("mysql", dsn)
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}

type User struct {
	Name string `fform:"PRIMARY KEY"`
	Age  int
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}

func transactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(session *session.Session) (result interface{}, err error) {
		// mysql ddl语句会自动进行一次commit
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 28})
		return nil, errors.New("error")
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}
