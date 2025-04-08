package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	api.Post("/connect", connectHandler)
	api.Post("/subscribe", subscribeHandler)

	if err := api.Listen("127.0.0.1:6080"); err != nil {
		log.Fatalln(err)
	}
}

// ConnectRequest ...
type ConnectRequest struct {
	Data RequestData `json:"data"`
}

// RequestData ...
type RequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConnectResponse ...
type ConnectResponse struct {
	Result     *ConnectResult `json:"result,omitempty"`
	Error      *Error         `json:"error,omitempty"`
	Disconnect *Disconnect    `json:"disconnect,omitempty"`
}

// ConnectResult ...
type ConnectResult struct {
	User string `json:"user"`
}

// Error ...
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// Disconnect ...
type Disconnect struct {
	Code   int    `json:"code"`
	Reason string `json:"reason,omitempty"`
}

// connectHandler ...
func connectHandler(ctx *fiber.Ctx) error {
	request := &ConnectRequest{}

	ctx.BodyParser(request)

	if request.Data.Username != "alex" || request.Data.Password != "qwerty" {
		return ctx.JSON(ConnectResponse{
			Error: &Error{
				Code:    101,
				Message: "unauthorized",
			},
		})
	}

	return ctx.JSON(ConnectResponse{
		Result: &ConnectResult{
			User: request.Data.Username,
		},
	})
}
