package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/jedib0t/go-pretty/v6/table"
)

var topologyPath = "topology/"

type topology struct {
	GenericResource
	Title   string `json:"title"`
	Doc     string `json:"documentation,omitempty"`
	Diagram string `json:"diagram_url,omitempty"`
}

type topoList struct {
	List []topology
}

func (c *Client) retrieveTopologies() ([]topology, error) {
	var result []topology

	resp, err := c.Get(simulationPath)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var topos topoList

	err = json.Unmarshal(body, &topos.List)
	if err != nil {
		return result, err
	}

	return topos.List, nil
}

func (c *Client) ListTopologies() error {

	topos, err := c.retrieveTopologies()
	if err != nil {
		return err
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "ID", "URL"})

	for _, topo := range topos {
		tw.AppendRow(table.Row{topo.Name, topo.Id, topo.Url})
	}

	fmt.Println(tw.Render())

	return nil
}

func topoGeneralizer(topos []topology) (result []GenericResourcer) {
	for _, topo := range topos {
		result = append(result, topo.GenericResource)
	}
	return result
}

func (c *Client) GetTopology(input string) error {

	topos, err := c.retrieveTopologies()
	if err != nil {
		return err
	}

	id, err := c.getResourceID(input, topoGeneralizer(topos))
	if err != nil {
		return err
	}

	resp, err := c.Get(topologyPath + id + "/")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var topo topology

	err = json.Unmarshal(body, &topo)
	if err != nil {
		return err
	}

	spew.Dump(topo)

	return nil
}
