package main

import (
	"s0611-23/internal/contacts"

	"github.com/gofiber/fiber/v2"
)

func main() {
	println("main")

	api := fiber.New()

	cont := contacts.NewHandler()

	api.Get("/api/v1/contact", cont.ContactRead)

}
