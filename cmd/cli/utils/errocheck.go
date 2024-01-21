package utils

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/grill/internal/grpcclient"
)

func CheckNoError(err error) {
	if err != nil {
		if errors.Is(err, err.(*grpcclient.UnauthenticatedError)) {
			log.Fatal("You must login first. use command: 'grill login'")
		} else {
			log.Fatal(err)
		}
	}
}
