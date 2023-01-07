package controllers

import (
	"context"
	"fmt"
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
	fmt.Println(code)

	token, err := githubcon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}
	fmt.Println(token)

	resp, err := http.Get("https://api.github.com/user/repo?access_token=" + token.AccessToken)
	//resp, err := http.Get('Authorization: token my_access_token' https://api.github.com/user/repos)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}
	fmt.Println(resp)

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}
	fmt.Println(userData)

	return c.SendString(string(userData))

}
