package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ServerFiber struct {
	app         *fiber.App
	errShutdown error
	port        uint
}

func NewFiber(portListening uint) *ServerFiber {
	return &ServerFiber{
		app:  fiber.New(),
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
		body := string(c.Body())
		fmt.Println("request body: ", body)

		return c.SendStatus(http.StatusOK)
	}
}

func (s *ServerFiber) handleGetCompany() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := string(c.Body())
		fmt.Println("request body: ", body)

		return c.SendStatus(http.StatusOK)
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
		body := string(c.Body())
		fmt.Println("request body: ", body)

		return c.SendStatus(http.StatusOK)
	}
}

func (s *ServerFiber) addRoutes() {
	s.app.Post(_route, s.handleNewCompany())
	s.app.Get(_route+"/:id", s.handleGetCompany())
	s.app.Get(_route, s.handleGetCompanies())
	s.app.Put(_route, s.handleUpdateCompany())
	s.app.Delete(_route, s.handleDeleteCompany())
}
