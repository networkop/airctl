package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var topologyPath = "topology/"

type topology struct {
	GenericResource
	Title   string `json:"title"`
	Diagram string `json:"diagram_url,omitempty"`
}

type TopoNotFound struct {
	Name string
}

type topoResponse struct {
	GenericResource
}

func (e *TopoNotFound) Error() string {
	return fmt.Sprintf("Topology file not found %s", e.Name)
}

func (c *Client) retrieveTopologies() ([]topology, error) {
	var result []topology

	resp, err := c.Get(topologyPath)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var topos []topology

	err = json.Unmarshal(body, &topos)
	if err != nil {
		return result, err
	}

	return topos, nil
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

func (c *Client) GetTopology(input string, quiet bool) error {

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

	if quiet {
		fmt.Println(topo.Id)
		return nil
	}

	transformer := text.NewJSONTransformer("", "\t")
	fmt.Println(transformer(topo))

	return nil
}

func (c *Client) CreateTopology(filename string) error {

	f, err := os.Open(filename)
	if os.IsNotExist(err) {
		return &TopoNotFound{Name: filename}
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Failed to read file %s", filename)
	}

	resp, err := c.PostDotFile(topologyPath, data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result topoResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	fmt.Println(result.Id)

	return nil
}

func (c *Client) DelTopology(inputs []string) error {

	topos, err := c.retrieveTopologies()
	if err != nil {
		return err
	}

	for _, input := range inputs {
		id, err := c.getResourceID(input, topoGeneralizer(topos))
		if err != nil {
			return err
		}

		_, err = c.Delete(topologyPath + id + "/")
		if err != nil {
			return err
		}

		fmt.Println(id)

	}

	return nil
}
