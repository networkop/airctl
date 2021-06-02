package air

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/networkop/airctl/internal/config"
	"github.com/sirupsen/logrus"
)

var (
	defaultBase     = "https://air.nvidia.com/api/v1/"
	jsonContentType = "application/json"
)

type LoginFailed struct {
	Err error
}

func (e *LoginFailed) Error() string {
	return fmt.Sprintf("Unauthenticated or login expired. Please re-login")
}

type Client struct {
	httpC *http.Client
	base  *url.URL
	token string
}

func NewClient() (*Client, error) {

	base, err := url.Parse(defaultBase)
	if err != nil {
		return nil, err
	}

	auth := config.NewAuthData()

	return &Client{
		httpC: &http.Client{
			Timeout: time.Second * 20,
		},
		base:  base,
		token: fmt.Sprintf("Bearer %s", auth.Token),
	}, nil
}
func (c *Client) Post(path string, payload interface{}) (*http.Response, error) {
	jsonReq, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.makePath(path), bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", jsonContentType)
	req.Header.Set("accept", jsonContentType)
	req.Header.Set("Authorization", c.token)

	logrus.Debugf("Request: \n%+v", req)
	resp, err := c.httpC.Do(req)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("Response: \n%+v", resp)
	return resp, nil
}

func (c *Client) Get(p string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, c.makePath(p), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", jsonContentType)
	req.Header.Set("accept", jsonContentType)
	req.Header.Set("Authorization", c.token)

	logrus.Debugf("Request: \n%+v", req)
	resp, err := c.httpC.Do(req)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("Response: \n%+v", resp)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return resp, err
	}

	if resp.StatusCode == 403 {
		return nil, &LoginFailed{}
	}

	return nil, fmt.Errorf("Received response %s", resp.Status)
}

func (c *Client) makePath(p string) string {

	path, _ := url.Parse(p)

	return c.base.ResolveReference(path).String()
}
