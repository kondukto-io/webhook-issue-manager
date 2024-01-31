package handler

import (
	"github.com/gofiber/fiber/v2"
)

type TestHandler interface {
	TestIssue(c *fiber.Ctx) error
}

type testHandler struct{}

func NewTestConnectionHandlers() TestHandler {
	return &testHandler{}
}

func (*testHandler) TestIssue(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusOK)
	c.JSON(fiber.Map{"connection": true})
	return nil
}
