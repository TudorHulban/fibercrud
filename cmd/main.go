package main

import (
	"fmt"
	"os"

	"github.com/TudorHulban/fibercrud/domain"
	"github.com/TudorHulban/fibercrud/infra"
	rest "github.com/TudorHulban/fibercrud/infra/http"
	repo "github.com/TudorHulban/fibercrud/repository"
)

func main() {
	infra.Initialization()

	repository, errNew := repo.NewRepoCompany()
	if errNew != nil {
		fmt.Printf("NewRepoCompany: %s", errNew.Error())
		os.Exit(1)
	}

	defer repository.DBConn.Close()

	repository.Migration(&domain.CompanyData{})

	fiber := rest.NewFiber(repository)
	defer fiber.Stop()

	fiber.Start()
}
