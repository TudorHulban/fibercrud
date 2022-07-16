package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type RepoCompany struct {
	DBConn *gorm.DB
}

func NewRepoCompany() (*RepoCompany, error) {
	return &RepoCompany{}, nil
}
