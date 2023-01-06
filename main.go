package main

import (
	"github.com/Siddheshk02/go-oauth2/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/google_login", controllers.GoogleLogin)
	//app.Post("/google_callback", controllers.GoogleCallback)
	app.All("/github_login", controllers.GithubLogin)
	//app.Post("/github_callback", controllers.GithubCallback)

	app.Listen(":8080")

}
