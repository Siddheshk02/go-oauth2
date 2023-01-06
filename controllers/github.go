package controllers

import (
	"github.com/Siddheshk02/go-oauth2/config"
	"github.com/gofiber/fiber/v2"
)

func GithubLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GitHubLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

/*func GithubCallback(c *fiber.Ctx) error {

}*/
