package cloud

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/libdyson-wg/opendyson/devices"
	"github.com/libdyson-wg/opendyson/internal/generated/oapi"
	"strings"
)

var printLn = fmt.Println

func resp2device(resp oapi.Device) devices.Device {
	base := devices.BaseDevice{
		Name:               resp.Name,
		ProductCategory:    string(resp.Category),
		Serial:             resp.SerialNumber,
		Model:              resp.Model,
		Type:               resp.Type,
		ConnectionCategory: devices.ConnectionCategory(resp.ConnectionCategory),
	}

	if resp.Variant != nil {
		base.Variant = *resp.Variant
	}

	if resp.ConnectedConfiguration != nil {
		connected := devices.BaseConnectedDevice{
			BaseDevice: base,
			Firmware: devices.Firmware{
				Version:             resp.ConnectedConfiguration.Firmware.Version,
				AutoUpdateEnabled:   resp.ConnectedConfiguration.Firmware.AutoUpdateEnabled,
				NewVersionAvailable: resp.ConnectedConfiguration.Firmware.NewVersionAvailable,
			},
		}

		connected.MQTT.TopicRoot = resp.ConnectedConfiguration.Mqtt.MqttRootTopicLevel
		connected.MQTT.Username = resp.SerialNumber
		connected.MQTT.Password = decryptPassword(resp.ConnectedConfiguration.Mqtt.LocalBrokerCredentials)

		var err error
		connected.IoT, err = GetDeviceIoT(base.Serial)
		if err != nil {
			_, _ = printLn(fmt.Sprintf("Error `%s` fetching IoT data for serial: %s", err, base.Serial))
		}
		return &connected
	}

	return &base
}

func mapIoT(resp oapi.IoTData) (iot devices.IoT) {
	iot.Endpoint = resp.Endpoint
	iot.TokenValue = resp.IoTCredentials.TokenValue
	iot.TokenKey = resp.IoTCredentials.TokenKey
	iot.TokenSignature = resp.IoTCredentials.TokenSignature
	iot.ClientID = resp.IoTCredentials.ClientId
	iot.CustomAuthorizerName = resp.IoTCredentials.CustomAuthorizerName

	return iot
}

func aesKey() []byte {
	return []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
}

func iv() []byte {
	return make([]byte, 16)
}

type passwordGrabber struct {
	Serial   string `json:"serial"`
	Password string `json:"apPasswordHash"`
}

func decryptPassword(in string) string {
	block, _ := aes.NewCipher(aesKey())
	bm := cipher.NewCBCDecrypter(block, iv())

	rawIn, _ := base64.StdEncoding.DecodeString(in)
	out := make([]byte, len(rawIn))
	bm.CryptBlocks(out, rawIn)

	out = []byte(strings.Trim(string(out), "\b"))

	grabber := passwordGrabber{}
	_ = json.Unmarshal(out, &grabber)

	return grabber.Password
}
