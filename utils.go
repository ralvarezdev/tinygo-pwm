package tinygo_pwm

// SetDuty sets the duty cycle of a PWM signal given the pulse width and period width in microseconds.
//
// Parameters:
//
// pwm: PWM to control its signal
// channel: PWM channel
// pulseMicroseconds: pulse width in microseconds
// periodMicroseconds: period width in microseconds
func SetDuty(pwm PWM, channel uint8, pulseMicroseconds, periodMicroseconds uint32) {
	// Avoid division by zero
	if periodMicroseconds == 0 {
		return
	}
	
	// Set the duty cycle
	if pwm != nil {
		pwm.Set(channel, uint32(float64(pwm.Top())*float64(pulseMicroseconds)/float64(periodMicroseconds)))
	}
}