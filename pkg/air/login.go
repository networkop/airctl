package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/networkop/airctl/internal/config"
)

var loginPath = "login/"

type LoginData struct {
	Id   string `json:"id"`
	Name string `json:"username"`
}

type LoginFailed struct {
}

func (e *LoginFailed) Error() string {
	return fmt.Sprintf("Login Failed")
}

func (c *Client) Login(token string) error {

	auth := config.NewAuthData()

	if err := auth.SaveAuth(token); err != nil {
		return err
	}

	client, err := NewClient()
	if err != nil {
		return err
	}
	resp, err := client.Get(loginPath)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var login LoginData

	err = json.Unmarshal(body, &login)
	if err != nil {
		return &LoginFailed{}
	}

	if login.Name == "" {
		return &LoginFailed{}
	}

	fmt.Println("Authentication successful")

	return nil
}
