package devices

type ConnectionCategory string

const (
	NonConnected ConnectionCategory = "nonConnected"
)

type Device interface {
	GetName() string
	GetSerial() string
	GetModel() string
	GetType() string
	GetVariant() string
	CanConnect() bool
}

type BaseDevice struct {
	Name               string             `yaml:"name"`
	Serial             string             `yaml:"serial"`
	Model              string             `yaml:"model"`
	Type               string             `yaml:"type"`
	Variant            string             `yaml:"variant"`
	ProductCategory    string             `yaml:"productCategory"`
	ConnectionCategory ConnectionCategory `yaml:"connectionCategory"`
}

func (b BaseDevice) GetName() string {
	return b.Name
}

func (b BaseDevice) GetSerial() string {
	return b.Serial
}

func (b BaseDevice) GetModel() string {
	return b.Model
}

func (b BaseDevice) GetType() string {
	return b.Type
}

func (b BaseDevice) GetVariant() string {
	return b.Variant
}

func (b BaseDevice) CanConnect() bool {
	return b.ConnectionCategory != NonConnected
}
