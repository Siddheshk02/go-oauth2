package controllers

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/Siddheshk02/go-oauth2/config"
	"github.com/gofiber/fiber/v2"
)

func GithubLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GitHubLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GithubCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	githubcon := config.GithubConfig()

	token, err := githubcon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://api.github.com/user/repos?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	return c.SendString(string(userData))

}
