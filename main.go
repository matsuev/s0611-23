package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	srv := fiber.New()

	srv.Post("/api/user", UserCreate)
	srv.Get("/api/user", UserRead)
	srv.Put("/api/user", UserUpdate)
	srv.Delete("/api/user", UserDelete)

	if err := srv.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}
