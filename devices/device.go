package devices

import (
	"github.com/google/uuid"
)

type Device struct {
	Name    string `yaml:"name"`
	Serial  string `yaml:"serial"`
	Model   string `yaml:"model"`
	Type    string `yaml:"type"`
	Variant string `yaml:"variant"`

	CloudIOT CloudIOTConfig `yaml:"cloud_iot"`
	MQTT     MQTTConfig     `yaml:"mqtt"`
	Firmware FirmwareData   `yaml:"firmware"`
}

type CloudIOTConfig struct {
	IoTCredentials CloudIOTCredentials `yaml:"iot_credentials"`
	Endpoint       string              `yaml:"endpoint"`
}

type CloudIOTCredentials struct {
	ClientID             uuid.UUID `yaml:"client_id"`
	CustomAuthorizerName string    `yaml:"custom_authorizer_name"`
	TokenKey             string    `yaml:"token_key"`
	TokenSignature       string    `yaml:"token_signature"`
	TokenValue           uuid.UUID `yaml:"token_value"`
}

type MQTTConfig struct {
	Port      int    `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	TopicRoot string `yaml:"topic_root"`
}

type FirmwareData struct {
	Version             string `yaml:"version"`
	AutoUpdateEnabled   bool   `yaml:"auto_update_enabled"`
	NewVersionAvailable bool   `yaml:"new_version_available"`
}
