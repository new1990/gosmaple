package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
     "github.com/new1990/gosmaple/entity"
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
    db, err = gorm.Open("postgres", "host=192.168.99.100 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
    if err != nil {
        panic(err)
    }

    autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

// Close is closing db
func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}

func autoMigration() {
       db.AutoMigrate(&entity.User{})
}
