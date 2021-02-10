package main

import (
	"errors"

	log "github.com/mikhno-s/zerologr"
)

func main() {
	fixByFile := []string{"asm_amd64.s", "proc.go", "icegatherer.go", "jsonrpc2"}
	fixByFunc := []string{"Handle"}

	log.Init("trace", fixByFile, fixByFunc)

	log.Infof("Hello %s", "friend")

	log.Errorf("%s", errors.New("This is an error message"))

	// log.Panicf("%s", errors.New("This is a panic"))
}
