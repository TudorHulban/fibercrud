package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TudorHulban/fibercrud/domain"
	"github.com/TudorHulban/fibercrud/infra"
	"github.com/TudorHulban/fibercrud/infra/auth"
	"github.com/TudorHulban/fibercrud/infra/pubk"
	repo "github.com/TudorHulban/fibercrud/repository"
	"github.com/gofiber/fiber/v2"
)

type ServerFiber struct {
	publisher infra.Publisher

	app  *fiber.App
	repo *repo.RepoCompany

	errShutdown error
	port        uint
}

const (
	_portFiber = 3000
	_route     = "/api/v1/company"
)

func NewFiber(repo *repo.RepoCompany) *ServerFiber {
	return &ServerFiber{
		app:       fiber.New(),
		repo:      repo,
		publisher: pubk.NewPublisherToKafka(),
		port:      _portFiber,
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

// handleNewCompany is handler for addition.
// manual testing: curl -X POST -H "Content-Type: application/json" --data "{\"code\": \"J1234\", \"name\": \"avata\", \"country\": \"Fidji\", \"website\": \"avata.fj\", \"phone\": \"+55 12345\"}" http://localhost:3000/api/v1/company
// manual testing: curl -X POST -H "Content-Type: application/json" --data "{\"code\": \"J5678\", \"name\": \"tommy\", \"country\": \"Tokelau\", \"website\": \"tommy.tk\", \"phone\": \"+25 5678\"}" http://localhost:3000/api/v1/company
func (s *ServerFiber) handleNewCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !c.IsFromLocal() {
			auth := auth.NewAuthorizerByIPApi()
			isAuthorized, errAuth := auth.IsAuthorized(c.IP())
			if errAuth != nil {
				return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
					"success": false,
					"error":   errAuth.Error(),
				})
			}

			if !isAuthorized {
				return c.Status(http.StatusForbidden).JSON(&fiber.Map{
					"success": false,
				})
			}
		}

		var data domain.CompanyData

		if errBody := c.BodyParser(&data); errBody != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errBody.Error(),
			})
		}

		if errValid := data.IsValid(); errValid != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errValid.Error(),
			})
		}

		company, errNew := domain.NewCompany(&data, s.repo)
		if errNew != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errNew.Error(),
			})
		}

		idInsert, errInsert := company.RepoNewCompany()
		if errInsert != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   errInsert.Error(),
			})
		}

		go s.publisher.PublishEvent(&company)

		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"id":      idInsert,
		})
	}
}

func (s *ServerFiber) handleGetCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idRequest := c.Params("id")
		idCompany, errReq := strconv.Atoi(idRequest)
		if errReq != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errReq.Error(),
			})
		}

		if idCompany < 1 {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   "company ID should at least 1",
			})
		}

		company, errNew := domain.NewCompanyEmpty(s.repo)
		if errNew != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   errNew.Error(),
			})
		}

		data, errGet := company.RepoGetCompanyByID(uint(idCompany))
		if errGet != nil {
			return c.Status(http.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"error":   errGet.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"company": data,
		})
	}
}

func (s *ServerFiber) handleGetCompanies() fiber.Handler {
	return func(c *fiber.Ctx) error {
		company, errNew := domain.NewCompanyEmpty(s.repo)
		if errNew != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   errNew.Error(),
			})
		}

		data, errSelect := company.RepoGetCompanies()
		if errSelect != nil {
			return c.Status(http.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"error":   errSelect.Error(),
			})
		}

		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"success":   true,
			"companies": data,
		})
	}
}

func (s *ServerFiber) handleUpdateCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !c.IsFromLocal() {
			auth := auth.NewAuthorizerByIPApi()
			isAuthorized, errAuth := auth.IsAuthorized(c.IP())
			if errAuth != nil {
				return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
					"success": false,
					"error":   errAuth.Error(),
				})
			}

			if !isAuthorized {
				return c.Status(http.StatusForbidden).JSON(&fiber.Map{
					"success": false,
				})
			}
		}

		var data domain.CompanyData

		if errBody := c.BodyParser(&data); errBody != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errBody.Error(),
			})
		}

		if errValid := data.IsValid(); errValid != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errValid.Error(),
			})
		}

		company, errNew := domain.NewCompany(&data, s.repo)
		if errNew != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errNew.Error(),
			})
		}

		go s.publisher.PublishEvent(&company)

		return c.Status(http.StatusOK).SendString(company.RepoUpdateCompany().Error() + "\n")
	}
}

func (s *ServerFiber) handleDeleteCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idRequest := c.Params("id")
		idCompany, errReq := strconv.Atoi(idRequest)
		if errReq != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   errReq.Error(),
			})
		}

		if idCompany < 1 {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   "company ID should at least 1",
			})
		}

		company, errNew := domain.NewCompanyEmpty(s.repo)
		if errNew != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   errNew.Error(),
			})
		}

		errDel := company.RepoDeleteCompany(uint(idCompany))
		if errDel != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   errDel.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
		})
	}
}

func (s *ServerFiber) addRoutes() {
	s.app.Post(_route, s.handleNewCompany())
	s.app.Get(_route+"/:id", s.handleGetCompany())
	s.app.Get(_route, s.handleGetCompanies())
	s.app.Put(_route, s.handleUpdateCompany())
	s.app.Delete(_route+"/:id", s.handleDeleteCompany())
}
