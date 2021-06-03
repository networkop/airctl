package air

import (
	"fmt"

	"github.com/google/uuid"
)

type GenericResourcer interface {
	GetName() string
	GetId() string
}

type GenericResource struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Url  string `json:"url"`
}

type Mylist struct {
	List []GenericResource
}

func (r GenericResource) GetName() string {
	return r.Name
}

func (r GenericResource) GetId() string {
	return r.Id
}

type MatchFailed struct {
	Match string
}

type MultipleMatch MatchFailed

func (e *MultipleMatch) Error() string {
	return fmt.Sprintf("Found more than one match for a name '%s'. Use UUID instead.\n", e.Match)
}

func (e *MatchFailed) Error() string {
	return fmt.Sprintf("Could not find a match for %s\n", e.Match)
}

// Gets resource ID trying to match input as UUID and Name
func (c *Client) getResourceID(input string, objs []GenericResourcer) (string, error) {
	_, err := uuid.Parse(input)
	if err == nil {
		return input, nil
	}

	var matches []string
	for _, obj := range objs {
		if obj.GetName() != "" && obj.GetName() == input {
			matches = append(matches, obj.GetId())
		}

	}

	if len(matches) > 1 {
		return "", &MultipleMatch{Match: input}
	}

	if len(matches) == 1 {
		return matches[0], nil
	}

	return "", &MatchFailed{Match: input}
}
