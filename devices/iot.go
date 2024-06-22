package devices

import (
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"net/http"
)

type IoT struct {
	Endpoint             string    `yaml:"endpoint"`
	ClientID             uuid.UUID `yaml:"client_id"`
	CustomAuthorizerName string    `yaml:"custom_authorizer_name"`
	TokenKey             string    `yaml:"token_key"`
	TokenSignature       string    `yaml:"token_signature"`
	TokenValue           uuid.UUID `yaml:"token_value"`
}

func (d *BaseConnectedDevice) iotOptions() (*paho.ClientOptions, error) {
	opts := paho.NewClientOptions()

	brokerAddress := fmt.Sprintf("wss://%s/mqtt", d.IoT.Endpoint)
	opts.AddBroker(brokerAddress)

	headers := http.Header{}
	for k, v := range map[string]string{
		d.IoT.TokenKey:                     d.IoT.TokenValue.String(),
		"X-Amz-CustomAuthorizer-Name":      d.CustomAuthorizerName,
		"X-Amz-CustomAuthorizer-Signature": d.TokenSignature,
	} {
		headers.Add(k, v)
	}

	opts.SetHTTPHeaders(headers)
	opts.SetClientID(d.IoT.ClientID.String())

	opts.SetProtocolVersion(3)

	return opts, nil
}
