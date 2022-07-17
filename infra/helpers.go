package infra

import (
	"os"

	repo "github.com/TudorHulban/fibercrud/repository"
)

func Initialization() {
	os.Remove(repo.RepoDBSQLite)
}
