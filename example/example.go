package main

import (
	"errors"

	"github.com/mikhno-s/zerologr"
)

func main() {
	// fixByFile := []string{"asm_amd64.s", "proc.go", "icegatherer.go", "jsonrpc2"}
	// fixByFunc := []string{"Handle"}
	log := zerologr.Init("trace")

	log.Infof("Hello %s", "friend")

	log.Errorf("%s", errors.New("This is an error message"))

	// log.Panicf("%s", err)
}
