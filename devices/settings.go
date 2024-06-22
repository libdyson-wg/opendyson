package devices

type Setting struct {
	key   SettingKey
	value string
}

type SettingKey string

const (
	FanPower             SettingKey = "fpwr"
	AutoMode             SettingKey = "auto"
	DirectionForward     SettingKey = "fdir"
	FanSpeed             SettingKey = "fnsp"
	NightMode            SettingKey = "nmod"
	SleepTimer           SettingKey = "sltm"
	ResetFilter          SettingKey = "rstf"
	ContinuousMonitoring SettingKey = "rhtm"
)
