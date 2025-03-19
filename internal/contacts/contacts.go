package contacts

import "github.com/gofiber/fiber/v2"

// Contact ...
type Contact struct{}

// Handler ...
type Handler struct{}

// NewHandler ...
func NewHandler() *Handler {
	return new(Handler)
}

// ContactRead ...
func (h *Handler) ContactRead(ctx *fiber.Ctx) error {
	return nil
}
