package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	authFile = "config"
	authDir  = ".air"
)

type AuthData struct {
	Token      string `json:"token"`
	configFile string
}

func NewAuthData() *AuthData {

	dirPath := filepath.Join(homeDir(), authDir)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			return nil
		}
	}

	filePath := filepath.Join(dirPath, authFile)
	return readAuth(filePath)
}

func readAuth(path string) *AuthData {
	result := &AuthData{
		configFile: path,
	}

	f, err := os.Open(result.configFile)
	if os.IsNotExist(err) {
		return result
	}

	raw, err := ioutil.ReadAll(f)
	if err != nil {
		logrus.Infof("Failed to read file %s: %s", path, err)
		return result
	}

	err = yaml.Unmarshal(raw, result)
	if err != nil {
		logrus.Infof("Failed to parse YAML %s: %s", path, err)
		return result
	}

	return result
}

func (c *AuthData) SaveAuth(token string) error {
	c.Token = token

	logrus.Debugf("Saving token data in %s", c.configFile)
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.configFile, bytes, 0600)
}

func homeDir() string {
	return os.Getenv("HOME")
}
