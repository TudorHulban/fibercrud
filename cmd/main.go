package main

import (
	"fibercrud/domain"
	"fibercrud/infra"
	rest "fibercrud/infra/http"
	repo "fibercrud/repository"
	"fmt"
	"os"
)

func main() {
	infra.Initialization()

	repo, errNew := repo.NewRepoCompany()
	if errNew != nil {
		fmt.Printf("NewRepoCompany: %s", errNew.Error())
		os.Exit(1)
	}

	defer repo.DBConn.Close()

	repo.Migration(&domain.CompanyData{})

	fiber := rest.NewFiber(repo)
	defer fiber.Stop()

	fiber.Start()
}
