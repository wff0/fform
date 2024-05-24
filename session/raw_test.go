package session

import (
	"database/sql"
	"fform/dialect"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var TestDB *sql.DB

const dsn = "root:123456@tcp(127.0.0.1:3306)/fform"

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("mysql", dsn)
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func New() *Session {
	d, _ := dialect.GetDialect("mysql")
	return NewSession(TestDB, d)
}

func TestSession_Exec(t *testing.T) {
	s := New()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2, but got", count)
	}
}

func TestSession_QueryRows(t *testing.T) {
	s := New()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	row := s.Raw("SELECT count(*) FROM User").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err)
	}
}
