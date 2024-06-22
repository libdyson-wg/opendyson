package main

import (
	"github.com/libdyson-wg/opendyson/cloud"
	"github.com/libdyson-wg/opendyson/cmd"
	"github.com/libdyson-wg/opendyson/internal/config"
)

func main() {
	tok, err := config.GetToken()
	if err != nil {
		panic(err)
	}
	cloud.SetToken(tok)
	cmd.Execute()
}
