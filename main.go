package main

import (
	"fmt"
	"os"
)

func initialization() {
	os.Remove(_repoDBSQLite)
}

func main() {
	initialization()

	repo, errNew := NewRepoCompany()
	if errNew != nil {
		fmt.Printf("NewRepoCompany: %s", errNew.Error())
		os.Exit(1)
	}

	defer repo.DBConn.Close()

	repo.Migration(&CompanyData{})

	fiber := NewFiber(_portFiber, repo)
	defer fiber.Stop()

	fiber.Start()
}
