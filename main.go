package main

import (
	"github.com/Siddheshk02/go-oauth2/config"
	"github.com/Siddheshk02/go-oauth2/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()
	config.GithubConfig()

	app.Get("/google_login", controllers.GoogleLogin)
	app.Get("/google_callback", controllers.GoogleCallback)
	app.Get("/github_login", controllers.GithubLogin)
	//app.Post("/github_callback", controllers.GithubCallback)

	app.Listen(":8080")

}
