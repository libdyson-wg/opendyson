package cloud

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/libdyson-wg/libdyson-go/internal/generated/oapi"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/libdyson-wg/libdyson-go/devices"
)

var (
	key = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	iv  = make([]byte, 16)
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

	out.MQTT.Password = decryptCredentials(in.ConnectedConfiguration.Mqtt.LocalBrokerCredentials)
	out.MQTT.TopicRoot = in.ConnectedConfiguration.Mqtt.MqttRootTopicLevel
	out.MQTT.Port = 1883
	out.MQTT.Username = in.SerialNumber

	out.Firmware.Version = in.ConnectedConfiguration.Firmware.Version
	out.Firmware.AutoUpdateEnabled = in.ConnectedConfiguration.Firmware.AutoUpdateEnabled
	out.Firmware.NewVersionAvailable = in.ConnectedConfiguration.Firmware.NewVersionAvailable
	return out
}

func decryptCredentials(in string) string {
	block, _ := aes.NewCipher(key)
	bm := cipher.NewCBCDecrypter(block, iv)

	rawIn, _ := base64.StdEncoding.DecodeString(in)
	out := make([]byte, len(rawIn))
	bm.CryptBlocks(out, rawIn)

	out = []byte(strings.Trim(string(out), "\b"))

	grabber := passwordGrabber{}
	_ = json.Unmarshal(out, &grabber)

	return grabber.Password
}

type passwordGrabber struct {
	Password string `json:"apPasswordHash"`
}
