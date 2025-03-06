package main

import "github.com/gofiber/fiber/v2"

// UserModel ...
type UserModel struct {
	ID       int
	Username string
}

// UserCreate ...
func UserCreate(ctx *fiber.Ctx) error {
	type Request struct {
		Username   string `json:"username"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
	}

	var req Request

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(req)
}

// UserRead ...
func UserRead(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")

	if id < 1 {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user := UserModel{
		ID:       id,
		Username: "alex",
	}

	return ctx.JSON(user)
}

// UserUpdate ...
func UserUpdate(ctx *fiber.Ctx) error {
	ctx.SendString("user update")

	return nil
}

// UserDelete ...
func UserDelete(ctx *fiber.Ctx) error {
	ctx.SendString("user delete")

	return nil
}
