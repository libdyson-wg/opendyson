package devices

import (
	"errors"
	"fmt"
)

type PureHotCool struct {
	*BaseConnectedDevice
}

func (d *PureHotCool) PowerOn() error {
	return d.Set(Setting{FanPower, "ON"})
}

func (d *PureHotCool) PowerOff() error {
	return d.Set(Setting{FanPower, "OFF"})
}

func (d *PureHotCool) AutoModeOn() error {
	return d.Set(Setting{AutoMode, "ON"})
}

func (d *PureHotCool) AutoModeOff() error {
	return d.Set(Setting{AutoMode, "OFF"})
}

func (d *PureHotCool) DirectionForward() error {
	return d.Set(Setting{DirectionForward, "ON"})
}

func (d *PureHotCool) DirectionReverse() error {
	return d.Set(Setting{DirectionForward, "OFF"})
}

func (d *PureHotCool) Speed(in int) error {
	if in < 0 || in > 10 {
		return errors.New("speed must be between 0 and 10")
	}

	return d.Set(
		Setting{FanPower, "ON"},
		Setting{FanSpeed, fmt.Sprintf("%04d", in)},
	)
}

func (d *PureHotCool) EnableContinuousMonitoring() error {
	return d.Set(Setting{ContinuousMonitoring, "ON"})
}

func (d *PureHotCool) DisableContinuousMonitoring() error {
	return d.Set(Setting{ContinuousMonitoring, "OFF"})
}

func (d *PureHotCool) EnableNightMode() error {
	return d.Set(Setting{NightMode, "ON"})
}

func (d *PureHotCool) DisableNightMode() error {
	return d.Set(Setting{NightMode, "OFF"})
}

func (d *PureHotCool) ResetFilter() error {
	return d.Set(Setting{ResetFilter, "RSTF"})
}
