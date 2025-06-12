package devices

import (
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/libdyson-wg/opendyson/internal/lan"
	"time"
)

type ConnectedMode int

const (
	ModeMQTT ConnectedMode = iota
	ModeIoT
)

var timeout = time.Second * 5

type ConnectedDevice interface {
	Set(in ...Setting) error
	Get(key ...SettingKey) (map[SettingKey]string, error)

	SendRaw(topic string, message []byte) error
	SubscribeRaw(topic string, handler func([]byte)) error
	ResolveLocalAddress() error

	CommandTopic() string
	StatusTopic() string
	FaultTopic() string

	SetMode(mode ConnectedMode)
}

type Firmware struct {
	Version             string `yaml:"version"`
	AutoUpdateEnabled   bool   `yaml:"auto_update_enabled"`
	NewVersionAvailable bool   `yaml:"new_version_available"`
}

var _ ConnectedDevice = new(BaseConnectedDevice)

type BaseConnectedDevice struct {
	BaseDevice

	MQTT     `yaml:"mqtt"`
	IoT      `yaml:"iot"`
	Firmware `yaml:"firmware"`

	mode   ConnectedMode
	client paho.Client
}

func (d *BaseConnectedDevice) initClient() error {
	if d.client != nil {
		return nil
	}

	var (
		opts *paho.ClientOptions
		err  error
	)

	switch d.mode {
	case ModeMQTT:
		opts, err = d.mqttOptions()
	case ModeIoT:
		opts, err = d.iotOptions()
	}

	if err != nil {
		return err
	}

	if opts == nil {
		return fmt.Errorf("mqtt options is nil")
	}

	c := paho.NewClient(opts)
	t := c.Connect()

	if !t.WaitTimeout(timeout) {
		return fmt.Errorf("mqtt connect %s timeout", d.MQTT.Address)
	}

	if t.Error() != nil {
		return fmt.Errorf("unable to connect: %w", t.Error())
	}

	d.client = c
	return nil
}

func (d *BaseConnectedDevice) Set(in ...Setting) error {
	return d.SendRaw(d.CommandTopic(), []byte{})
}

func (d *BaseConnectedDevice) Get(key ...SettingKey) (map[SettingKey]string, error) {
	return nil, nil
}

func (d *BaseConnectedDevice) SetMode(mode ConnectedMode) {
	if d.mode == mode {
		return
	}
	if d.client != nil {
		d.client.Disconnect(250)
		d.client = nil
	}
	d.mode = mode
}

func (d *BaseConnectedDevice) SendRaw(topic string, msg []byte) error {
	if err := d.initClient(); err != nil {
		return err
	}

	qos := byte(2)
	if d.mode == ModeIoT {
		qos = 1
	}

	t := d.client.Publish(topic, qos, false, msg)
	if !t.WaitTimeout(timeout) {
		return fmt.Errorf("timeout sending message")
	}
	return t.Error()
}

func (d *BaseConnectedDevice) SubscribeRaw(topic string, callback func([]byte)) error {
	if err := d.initClient(); err != nil {
		return err
	}

	t := d.client.Subscribe(topic, 0, func(client paho.Client, msg paho.Message) {
		msg.Ack()
		callback(msg.Payload())
	})
	if !t.WaitTimeout(timeout) {
		return fmt.Errorf("timeout subscribing to topic %s", topic)
	}
	return t.Error()
}

func (d *BaseConnectedDevice) CommandTopic() string {
	return fmt.Sprintf("%s/%s/command", d.MQTT.TopicRoot, d.Serial)
}

func (d *BaseConnectedDevice) StatusTopic() string {
	return fmt.Sprintf("%s/%s/status/current", d.MQTT.TopicRoot, d.Serial)
}

func (d *BaseConnectedDevice) FaultTopic() string {
	return fmt.Sprintf("%s/%s/status/fault", d.MQTT.TopicRoot, d.Serial)
}

func (d *BaseConnectedDevice) ResolveLocalAddress() error {
	select {
	case ip := <-lan.RequestAddress(d.Serial):
		if ip == nil {
			return fmt.Errorf("unable to get ip address")
		}

		d.MQTT.Address = ip.String()
	}

	return nil
}
