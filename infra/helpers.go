package infra

import (
	repo "fibercrud/repository"
	"os"
)

func Initialization() {
	os.Remove(repo.RepoDBSQLite)
}
