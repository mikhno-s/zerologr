package main

import (
	"errors"

	log "github.com/mikhno-s/zerologr"
)

func init() {
	// fixByFile := []string{"asm_amd64.s", "proc.go"}
	// fixByFunc := []string{}
}

func main() {
	fixByFile := []string{"asm_amd64.s", "proc.go", "icegatherer.go", "jsonrpc2"}
	fixByFunc := []string{"Handle"}
	log.Init("error", fixByFile, fixByFunc)

	log.Infof("Hello %s", "friend")

	err := errors.New("This is an error message")
	log.Errorf("%s", err)

	l()
}

func l() {
	err := errors.New("This is an error message")
	log.Errorf("%s", err)
}
