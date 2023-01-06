package controllers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GithubLogin(c *fiber.Ctx) error {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	var GitHubLoginConfig oauth2.Config = oauth2.Config{
		RedirectURL:  "http://localhost:8080/github_callback",
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"name", "email", "repo"},
		Endpoint:     github.Endpoint,
	}

	url := GitHubLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

/*func GithubCallback(c *fiber.Ctx) error {

}*/
