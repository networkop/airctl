package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/sirupsen/logrus"
)

var servicePath = "service/"

type service struct {
	GenericResource
	Sim  string `json:"simulation"`
	Type string `json:"service_type"`
	Link string `json:"link"`
}

type serviceRequest struct {
	Sim   string `json:"simulation"`
	Intf  string `json:"interface"`
	Name  string `json:"name"`
	Type  string `json:"service_type"`
	DPort int    `json:"dest_port"`
}

func (c *Client) retrieveServices() ([]service, error) {
	var result []service

	resp, err := c.Get(servicePath)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	logrus.Debugf("GOT payload: %+v", string(body))

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Retrieves the list of all services and prints them in a table
func (c *Client) ListServices(quiet bool) error {

	svcs, err := c.retrieveServices()
	if err != nil {
		return err
	}

	if quiet {
		for _, svc := range svcs {
			fmt.Println(svc.Id)
		}
		return nil
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "Id", "Sim", "Type", "Link"})

	for _, svc := range svcs {
		tw.AppendRow(table.Row{svc.Name, svc.Id, svc.Sim, svc.Type, svc.Link})
	}

	fmt.Println(tw.Render())

	return nil
}

func serviceGeneralizer(svcs []service) (result []GenericResourcer) {
	for _, svc := range svcs {
		result = append(result, svc.GenericResource)
	}
	return result
}

// Prints the information about a specific simulation
func (c *Client) GetService(input string, quiet bool) error {

	svcs, err := c.retrieveServices()
	if err != nil {
		return err
	}

	id, err := c.getResourceID(input, serviceGeneralizer(svcs))
	if err != nil {
		return err
	}

	resp, err := c.Get(servicePath + id + "/")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var svc service

	err = json.Unmarshal(body, &svc)
	if err != nil {
		return err
	}

	if quiet {
		fmt.Println(svc.Id)
		return nil
	}

	transformer := text.NewJSONTransformer("", "\t")
	fmt.Println(transformer(svc))

	return nil
}

// Prints the information about a specific simulation
func (c *Client) DelService(inputs []string) error {

	svcs, err := c.retrieveServices()
	if err != nil {
		return err
	}

	for _, input := range inputs {
		id, err := c.getResourceID(input, serviceGeneralizer(svcs))
		if err != nil {
			return err
		}

		_, err = c.Delete(servicePath + id + "/")
		if err != nil {
			return err
		}

		fmt.Println(id)

	}

	return nil
}

func (c *Client) CreateSSHService(intf string, simID, name string) error {

	intfID, err := c.resolveInterface(intf, simID)
	if err != nil {
		return err
	}
	logrus.Debugf("Identified interface id : %s", intfID)

	if name == "" {
		name = fmt.Sprintf(("%s-ssh"), intf)
	}

	resp, err := c.Post(servicePath, serviceRequest{
		Name:  name,
		Sim:   simID,
		Intf:  intfID,
		Type:  "ssh",
		DPort: 22,
	})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var svc service

	err = json.Unmarshal(body, &svc)
	if err != nil {
		return err
	}

	var result GenericResource
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	fmt.Println(result.Id)
	return nil
}
