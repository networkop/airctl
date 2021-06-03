package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var nodePath = "node/"

type InterfaceResolve struct {
	Intf string
}

type Node struct {
	Name  string      `json:"name"`
	Intfs []Interface `json:"interfaces"`
}

type Interface struct {
	GenericResource
}

func (e *InterfaceResolve) Error() string {
	return fmt.Sprintf("Unable to resolve interface %s", e.Intf)
}

func (c *Client) resolveInterface(intf, simId string) (string, error) {

	parts := strings.Split(intf, ":")
	if len(parts) != 2 {
		return "", &InterfaceResolve{Intf: intf}
	}

	resp, err := c.Get(nodePath + fmt.Sprintf("?simulation=%s", simId))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []Node

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	nodeName, intfName := parts[0], parts[1]
	for _, node := range result {
		if node.Name != nodeName {
			continue
		}

		for _, i := range node.Intfs {
			if i.Name == intfName {
				return c.getSimInterface(i.Id, simId)
			}
		}
	}

	return "", &InterfaceResolve{Intf: intf}
}
