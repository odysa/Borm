package Borm

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestNewEngine(t *testing.T) {
	db,err := NewEngine("sqlite3","test.db")
	if err!=nil{
		t.Error("Failed to connect database")
	}
	s:=db.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
