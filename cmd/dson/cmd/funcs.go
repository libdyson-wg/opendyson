package cmd

import (
	"github.com/libdyson-wg/libdyson-go/cloud"
	"github.com/libdyson-wg/libdyson-go/config"
	"github.com/libdyson-wg/libdyson-go/devices"
	"github.com/libdyson-wg/libdyson-go/internal/account"
	"github.com/libdyson-wg/libdyson-go/internal/shell"
)

type functions struct {
	Login      func() error
	GetDevices func() ([]devices.Device, error)
}

var funcs functions

func init() {
	funcs = functions{
		Login: account.Login(
			shell.PromptForInput,
			shell.PromptForPassword,
			cloud.BeginLogin,
			cloud.CompleteLogin,
			config.SetToken,
			cloud.SetToken,
			cloud.SetServerRegion,
		),
		GetDevices: cloud.GetDevices,
	}

}
