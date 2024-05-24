package fform

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/ffgorm"

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
