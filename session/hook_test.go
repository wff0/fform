package session

import (
	"fform/log"
	"testing"
)

type Account struct {
	ID       int `fform:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	account.ID += 1000
	log.Info("before inert", account)
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	account.Password = "******"
	log.Info("after query", account)
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := New().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})

	u := &Account{}

	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", u)
	}
}
