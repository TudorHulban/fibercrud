package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ServerFiber struct {
	app  *fiber.App
	repo *RepoCompany

	errShutdown error
	port        uint
}

func NewFiber(portListening uint, repo *RepoCompany) *ServerFiber {
	return &ServerFiber{
		app:  fiber.New(),
		repo: repo,
		port: portListening,
	}
}

func (s *ServerFiber) Start() {
	s.addRoutes()

	s.errShutdown = s.app.Listen(":" + strconv.Itoa(int(s.port)))
}

func (s *ServerFiber) Stop() {
	fmt.Println("stopping Fiber")

	if errShut := s.app.Shutdown(); errShut != nil {
		fmt.Printf("error Fiber: %s\n", errShut.Error())
	}
}

func (s *ServerFiber) handleNewCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := CompanyData{
			Code:    "xyz",
			Name:    "Ltd",
			Country: "UK",
			Website: "dice.com",
			Phone:   "+40 123 456",
		}

		if errValid := data.IsValid(); errValid != nil {
			return c.Status(http.StatusBadRequest).Send([]byte(errValid.Error() + "\n"))
		}

		company, errNew := NewCompany(&data, s.repo)
		if errNew != nil {
			return c.Status(http.StatusBadRequest).Send([]byte(errNew.Error() + "\n"))
		}

		company.RepoNewCompany()

		return c.SendStatus(http.StatusOK)
	}
}

func (s *ServerFiber) handleGetCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idRequest := c.Params("id")
		idCompany, errReq := strconv.Atoi(idRequest)
		if errReq != nil {
			return c.Status(http.StatusBadRequest).Send([]byte(errReq.Error() + "\n"))
		}

		if idCompany < 1 {
			return c.Status(http.StatusBadRequest).Send([]byte("company ID should at least 1" + "\n"))
		}

		company, errNew := NewCompanyEmpty(s.repo)
		if errNew != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte(errNew.Error() + "\n"))
		}

		data, errGet := company.RepoGetCompany(uint(idCompany))
		if errGet != nil {
			return c.Status(http.StatusNotFound).Send([]byte(errGet.Error() + "\n"))
		}

		return c.JSON(data)
	}
}

func (s *ServerFiber) handleGetCompanies() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := string(c.Body())
		fmt.Println("request body: ", body)

		return c.SendStatus(http.StatusOK)
	}
}

func (s *ServerFiber) handleUpdateCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := string(c.Body())
		fmt.Println("request body: ", body)

		return c.SendStatus(http.StatusOK)
	}
}

func (s *ServerFiber) handleDeleteCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idRequest := c.Params("id")
		idCompany, errReq := strconv.Atoi(idRequest)
		if errReq != nil {
			return c.Status(http.StatusBadRequest).Send([]byte(errReq.Error() + "\n"))
		}

		if idCompany < 1 {
			return c.Status(http.StatusBadRequest).Send([]byte("company ID should at least 1" + "\n"))
		}

		company, errNew := NewCompanyEmpty(s.repo)
		if errNew != nil {
			return c.Status(http.StatusInternalServerError).Send([]byte(errNew.Error() + "\n"))
		}

		errDel := company.RepoDeleteCompany(uint(idCompany))
		if errDel != nil {
			return c.Status(http.StatusOK).Send([]byte(errDel.Error() + "\n"))
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (s *ServerFiber) addRoutes() {
	s.app.Post(_route, s.handleNewCompany())
	s.app.Get(_route+"/:id", s.handleGetCompany())
	s.app.Get(_route, s.handleGetCompanies())
	s.app.Put(_route, s.handleUpdateCompany())
	s.app.Delete(_route+"/:id", s.handleDeleteCompany())
}
