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
	log.Init()
	log.Infof("Hello %s", "friend")

	err := errors.New("This is an error message")
	log.Errorf("debug server stopped. got err: %s", err)
}
