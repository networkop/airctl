package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	simulationPath = "simulation/"
	citcPath       = "citc/"

	SimulationState = struct {
		Up, Down, Destroy, Duplicate string
	}{
		Up:        "up",
		Down:      "down",
		Destroy:   "destroy",
		Duplicate: "duplicate",
	}
)

type simulation struct {
	GenericResource
	State    string   `json:"state"`
	Title    string   `json:"title"`
	Services []string `json:"services"`
}

type simList struct {
	List []simulation
}

type simAction struct {
	Action string `json:"action"`
	Name   string `json:"name,omitempty"`
	Title  string `json:"title,omitempty"`
}

type simResult struct {
	Result     string     `json:"result"`
	Simulation simulation `json:"simulation"`
}

func (c *Client) retrieveSimulations() ([]simulation, error) {
	var result simList

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

// Retrieves the list of all simulations and prints them in a table
func (c *Client) ListSimulations(quiet bool) error {

	sims, err := c.retrieveSimulations()
	if err != nil {
		return err
	}

	if quiet {
		for _, sim := range sims {
			fmt.Println(sim.Id)
		}
		return nil
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "ID", "State"})

	for _, sim := range sims {
		tw.AppendRow(table.Row{sim.Name, sim.Id, sim.State})
	}

	fmt.Println(tw.Render())

	return nil
}

func (c *Client) retrieveSimulation(id string) (simulation, error) {
	var result simulation

	resp, err := c.Get(simulationPath + id + "/")
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func simGeneralizer(sims []simulation) (result []GenericResourcer) {
	for _, sim := range sims {
		result = append(result, sim.GenericResource)
	}
	return result
}

// Prints the information about a specific simulation
func (c *Client) GetSimulation(input string, quiet bool) error {

	sims, err := c.retrieveSimulations()
	if err != nil {
		return err
	}

	id, err := c.getResourceID(input, simGeneralizer(sims))
	if err != nil {
		return err
	}

	sim, err := c.retrieveSimulation(id)
	if err != nil {
		return err
	}

	if quiet {
		fmt.Println(sim.Id)
		return nil
	}

	transformer := text.NewJSONTransformer("", "\t")
	fmt.Println(transformer(sim))

	return nil
}

func (c *Client) SetSimulation(id string, state string) error {
	sim, err := c.retrieveSimulation(id)
	if err != nil {
		return err
	}

	switch state {
	case SimulationState.Up:
		if sim.State != "LOADED" {
			_, err := c.Post(simulationPath+id+"/control/", simAction{Action: "load"})
			if err != nil {
				return err
			}
		}
	case SimulationState.Down:
		if sim.State == "LOADED" {
			_, err := c.Post(simulationPath+id+"/control/", simAction{Action: "store"})
			if err != nil {
				return err
			}
		}
	case SimulationState.Destroy:
		_, err := c.Post(simulationPath+id+"/control/", simAction{Action: SimulationState.Destroy})
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Unrecognised state: %s", state)

	}

	return nil
}

func (c *Client) CreateSimulation(simID string) error {
	_, err := uuid.Parse(simID)
	if err != nil {
		return fmt.Errorf("Malformed simID: %s", simID)
	}

	resp, err := c.Post(simulationPath+simID+"/control/", simAction{
		Action: SimulationState.Duplicate,
	})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result simResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if result.Result != "success" {
		return fmt.Errorf("failed to create sim %s", result.Result)
	}

	fmt.Println(result.Simulation.Id)

	return nil
}

func (c *Client) retrieveCITC() (simulation, error) {
	var result simulation

	resp, err := c.Get(simulationPath + citcPath)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *Client) CreateCITC() error {

	citc, err := c.retrieveCITC()
	if err != nil {
		return err
	}

	return c.CreateSimulation(citc.Id)
}
