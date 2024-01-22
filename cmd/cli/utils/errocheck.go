package utils

import (
	"github.com/charmbracelet/log"
	"strings"
)

func CheckNoError(err error) {
	if err != nil {
		// check if the error is a grpc error with a status Unauthenticated
		if strings.Contains(err.Error(), "Unauthenticated") {
			log.Fatal("You must login first. use command: 'grill login'")
		} else {
			log.Fatal(err)
		}
	}
}
