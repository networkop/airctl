package air

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var accountPath = "account/"

type account struct {
	Url       string `json:"url"`
	ID        string `json:"id"`
	Worker    string `json:"worker,omitempty"`
	Name      string `json:"username,omitempty"`
	Admin     bool   `json:"admin,omitempty"`
	Staff     bool   `json:"staff,omitempty"`
	LastLogin string `json:"last_login,omitempty"`
}

type accList struct {
	List []account
}

func (c *Client) retrieveAccounts() ([]account, error) {
	var result accList

	resp, err := c.Get(accountPath)
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

func (c *Client) ListAccounts(quiet bool) error {

	accs, err := c.retrieveAccounts()
	if err != nil {
		return err
	}

	if quiet {
		for _, acc := range accs {
			fmt.Println(acc.ID)
		}
		return nil
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Id", "Username", "Admin", "Worker", "Last Login"})

	for _, acc := range accs {
		tw.AppendRow(table.Row{acc.ID, acc.Name, acc.Admin, acc.Worker, acc.LastLogin})
	}
	tw.AppendFooter(table.Row{"", "", "Total", len(accs)})

	fmt.Println(tw.Render())

	return nil
}

// we can't use generic resource function because there's no `name` in account
func (c *Client) getAccountID(input string) string {

	_, err := uuid.Parse(input)
	if err == nil {
		return input
	}

	accs, err := c.retrieveAccounts()
	if err != nil {
		return ""
	}

	for _, resource := range accs {

		if resource.Name != "" && resource.Name == input {
			return resource.ID
		}

	}

	return ""

}

func (c *Client) GetAccount(input string, quiet bool) error {

	id := c.getAccountID(input)

	resp, err := c.Get(accountPath + id + "/")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var acc account

	err = json.Unmarshal(body, &acc)
	if err != nil {
		return err
	}

	if quiet {
		fmt.Println(acc.ID)
		return nil
	}

	transformer := text.NewJSONTransformer("", "\t")
	fmt.Println(transformer(acc))

	return nil
}
