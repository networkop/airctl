package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var servicePath = "service/"

type service struct {
	Url   string `json:"url"`
	Id    string `json:"id"`
	Name  string `json:"name"`
	Sim   string `json:"simulation"`
	Type  string `json:"service_type"`
	DPort int    `json:"src_port"`
	Host  string `json:"host"`
}

type servList struct {
	List []service
}

func (c *Client) retrieveServices() ([]service, error) {
	var result servList

	resp, err := c.Get(simulationPath)
	if err != nil {
		return result.List, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result.List, err
	}

	err = json.Unmarshal(body, &result.List)
	if err != nil {
		return result.List, err
	}
	return result.List, nil
}

// Retrieves the list of all services and prints them in a table
func (c *Client) ListServices() error {

	svcs, err := c.retrieveServices()
	if err != nil {
		return err
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "Id", "Sim", "Type", "Host", "Port"})

	for _, svc := range svcs {
		tw.AppendRow(table.Row{svc.Name, svc.Id, svc.Sim, svc.Type, svc.Host, svc.DPort})
	}

	fmt.Println(tw.Render())

	return nil
}

// Prints the information about a specific simulation
func (c *Client) GetService(input string) error {

	svcs, err := c.retrieveServices()
	if err != nil {
		return err
	}

	id, err := c.getResourceID(input, svcs)
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

	transformer := text.NewJSONTransformer("", "\t")
	fmt.Println(transformer(svc))

	return nil
}
