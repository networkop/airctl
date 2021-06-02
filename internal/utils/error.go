package utils

import (
	"fmt"

	"github.com/networkop/airctl/pkg/air"
)

func ProcessError(err error) error {
	_, ok := err.(*air.LoginFailed)
	if ok {
		fmt.Println(err.Error())
		return nil
	}

	return err
}
