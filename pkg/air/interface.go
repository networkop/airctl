package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var simInterfacePath = "simulation-interface/"

type simInterface struct {
	Id       string `json:"id"`
	Original string `json:"original"`
}

func (c *Client) getSimInterface(intfId, simId string) (string, error) {

	resp, err := c.Get(simInterfacePath + fmt.Sprintf("?simulation=%s", simId))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []simInterface

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	for _, intf := range result {
		if strings.Contains(intf.Original, intfId) {
			return intf.Id, nil
		}

	}

	return "", &InterfaceResolve{Intf: intfId}
}
