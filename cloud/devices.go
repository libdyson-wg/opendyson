package cloud

import (
	"context"
	"fmt"
	"github.com/libdyson-wg/libdyson-go/internal/generated/oapi"
	"net/http"
	"os"
	"time"

	"github.com/libdyson-wg/libdyson-go/devices"
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
		ds[i] = mapDevice((*resp.JSON200)[i])

		resp, err := client.GetIoTInfoWithResponse(ctx, oapi.GetIoTInfoJSONRequestBody{
			Serial: ds[i].Serial,
		})

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "error getting IoTInfo from cloud for", ds[i].Serial)
			continue
		}

		ds[i].CloudIOT = mapIoT(*resp.JSON200)
	}

	return ds, nil
}

func mapIoT(in oapi.IoTData) (out devices.CloudIOTConfig) {
	out.Endpoint = in.Endpoint
	out.IoTCredentials.CustomAuthorizerName = in.IoTCredentials.CustomAuthorizerName
	out.IoTCredentials.ClientID = in.IoTCredentials.ClientId
	out.IoTCredentials.TokenKey = in.IoTCredentials.TokenKey
	out.IoTCredentials.TokenSignature = in.IoTCredentials.TokenSignature
	out.IoTCredentials.TokenValue = in.IoTCredentials.TokenValue
	return out
}

func mapDevice(in oapi.Device) (out devices.Device) {
	out.Model = in.Model
	out.Name = in.Name
	out.Serial = in.SerialNumber
	out.Type = in.Type
	if in.Variant != nil {
		out.Variant = *in.Variant
	}

	out.MQTT.LocalCredentials = in.ConnectedConfiguration.Mqtt.LocalBrokerCredentials
	out.MQTT.TopicRoot = in.ConnectedConfiguration.Mqtt.MqttRootTopicLevel

	out.Firmware.Version = in.ConnectedConfiguration.Firmware.Version
	out.Firmware.AutoUpdateEnabled = in.ConnectedConfiguration.Firmware.AutoUpdateEnabled
	out.Firmware.NewVersionAvailable = in.ConnectedConfiguration.Firmware.NewVersionAvailable
	return out
}
