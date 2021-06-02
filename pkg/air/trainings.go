package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jedib0t/go-pretty/v6/table"
)

var trainingPath = "../../training/api/v1/training/"

type training struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Sim    string `json:"simulation_id"`
	Public bool   `json:"public"`
}

type trainingList struct {
	List []training `json:"results"`
}

func (c *Client) GetTrainings() error {

	resp, err := c.Get(trainingPath)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var trainings trainingList

	err = json.Unmarshal(body, &trainings)
	if err != nil {
		return err
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "Sim ID", "Public"})

	for _, training := range trainings.List {
		tw.AppendRow(table.Row{training.Name, training.Sim, training.Public})
	}

	fmt.Println(tw.Render())

	return nil
}
