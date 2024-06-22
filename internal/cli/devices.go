package cli

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/libdyson-wg/opendyson/devices"
)

func DeviceGetter(getDevices func() ([]devices.Device, error)) func() ([]devices.Device, error) {
	return func() ([]devices.Device, error) {
		ds, err := getDevices()
		if err != nil {
			return nil, err
		}

		wg := sync.WaitGroup{}
		for _, d := range ds {
			if cd, ok := d.(devices.ConnectedDevice); ok {
				wg.Add(1)
				go func(cd devices.ConnectedDevice) {
					err := cd.ResolveLocalAddress()
					if err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}(cd)
			}
		}
		wg.Wait()
		return ds, nil
	}
}

func Listener(
	getDevices func() ([]devices.Device, error),
	printLine func(in string),
) func(serial string, iot bool) error {
	return func(serial string, iot bool) error {
		ds, err := getDevices()
		if err != nil {
			return err
		}

		var d devices.Device

		for i := range ds {
			if ds[i].GetSerial() == serial {
				d = ds[i]
			}
		}

		if d == nil {
			return fmt.Errorf("device with serial %s not found", serial)
		}

		var (
			cd devices.ConnectedDevice
			ok bool
		)

		if cd, ok = d.(devices.ConnectedDevice); !ok {
			return fmt.Errorf("device %s is not connected", serial)
		}

		if iot {
			cd.SetMode(devices.ModeIoT)
		}

		for name, topic := range map[string]string{
			"Status:   ": cd.StatusTopic(),
			"Fault:    ": cd.FaultTopic(),
			"Command:  ": cd.CommandTopic(),
		} {
			printLine(fmt.Sprintf("Subscribing to %s", topic))
			if err = cd.SubscribeRaw(topic, func(bytes []byte) {
				printLine(fmt.Sprintf("%s%s", name, string(bytes)))
			}); err != nil {
				return err
			}
		}

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM, os.Interrupt)
		go func() {
			<-sig
			os.Exit(0)
		}()

		select {}
	}
}
