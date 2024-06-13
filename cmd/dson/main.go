package main

import (
	"github.com/libdyson-wg/libdyson-go/cloud"
	"github.com/libdyson-wg/libdyson-go/cmd/dson/cmd"
	"github.com/libdyson-wg/libdyson-go/config"
)

func main() {
	tok, err := config.GetToken()
	if err != nil {
		panic(err)
	}
	cloud.SetToken(tok)
	cmd.Execute()
}
