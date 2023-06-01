package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDb struct {
	db *gorm.DB
}

func (ctx *SqliteDb) InitializeDatabase() {
	db, err := gorm.Open(sqlite.Open("/Users/harish/Work/learning-golang/todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ctx.db = db
	ctx.db.AutoMigrate(&Todo{})
}

func (ctx *SqliteDb) GetAllTodo() ([]Todo, error) {
	var todos []Todo
	ctx.db.Find(&todos)
	return todos, nil
}
