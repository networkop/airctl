package air

import (
	"github.com/networkop/airctl/internal/config"
	"github.com/sirupsen/logrus"
)

var loginPath = "login/"

func Login(token string) error {

	logrus.Infof("Authenticating")
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

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		logrus.Infof("Authentication successful")

	}

	return nil
}
