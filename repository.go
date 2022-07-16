package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type RepoCompany struct {
	DBConn *gorm.DB
}

func NewRepoCompany() (*RepoCompany, error) {
	dbConn, errOpen := gorm.Open("sqlite3", _repoDBSQLite)
	if errOpen != nil {
		return nil, fmt.Errorf("NewRepoCompany gorm.Open: %w", errOpen)
	}

	return &RepoCompany{
		DBConn: dbConn,
	}, nil
}

func (r *RepoCompany) Migration(model any) {
	r.DBConn.AutoMigrate(model)
}
