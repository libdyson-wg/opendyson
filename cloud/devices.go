package cloud

import (
	"context"
	"fmt"
	"github.com/libdyson-wg/opendyson/internal/generated/oapi"
	"net/http"
	"time"

	"github.com/libdyson-wg/opendyson/devices"
)

func GetDevices() ([]devices.Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	resp, err := client.GetDevicesWithResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting devices from cloud: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("error getting devices from cloud, http status code: %d", resp.StatusCode())
	}

	ds := make([]devices.Device, len(*resp.JSON200))
	for i := 0; i < len(ds); i++ {
		ds[i] = resp2device((*resp.JSON200)[i])
	}

	return ds, nil
}

func GetDeviceIoT(serial string) (devices.IoT, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.GetIoTInfoWithResponse(ctx, oapi.GetIoTInfoJSONRequestBody{
		Serial: serial,
	})

	if err != nil {
		return devices.IoT{}, fmt.Errorf("error getting IoT info from cloud for %s, %w", serial, err)
	}

	if resp.StatusCode() != http.StatusOK {
		return devices.IoT{}, fmt.Errorf("error getting IoT info from cloud, http status code: %d", resp.StatusCode())
	}

	return mapIoT(*resp.JSON200), nil
}
