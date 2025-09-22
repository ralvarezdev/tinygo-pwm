package tinygo_pwm

import (
	"machine"
)

type (
	// PWM is the interface necessary for controlling PWM signals.
	PWM interface {
		Configure(config machine.PWMConfig) error
		Channel(pin machine.Pin) (channel uint8, err error)
		Top() uint32
		Set(channel uint8, value uint32)
	}
)
