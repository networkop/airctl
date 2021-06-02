package air

import (
	"fmt"

	"github.com/google/uuid"
)

type GenericResource struct {
	Name string
	Id   string
}

type GenericResourceList struct {
	List []GenericResource
}

// Gets resource ID trying to match input as UUID and Name
func (c *Client) getResourceID(input string, objs interface{}) (string, error) {

	_, err := uuid.Parse(input)
	if err == nil {
		return input, nil
	}

	l, ok := objs.(GenericResourceList)
	if !ok {
		return "", fmt.Errorf("Malformed resource list %+v", objs)
	}

	var matches []string
	for _, resource := range l.List {

		if resource.Name != "" && resource.Name == input {
			matches = append(matches, resource.Id)
		}

	}

	if len(matches) > 1 {
		return "", fmt.Errorf("Found more than one match with name %s", input)
	}

	return "", fmt.Errorf("Could not find a match for %s", input)
}
