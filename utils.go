package tinygo_pwm

// SetDuty sets the duty cycle of a PWM signal given the pulse width and period width.
//
// Parameters:
//
// pwm: PWM to control its signal
// channel: PWM channel
// pulse: pulse width
// period: period width
func SetDuty(pwm PWM, channel uint8, pulse, period uint32) {
	// Avoid division by zero
	if period == 0 {
		return
	}
	
	// Set the duty cycle
	if pwm != nil {
		pwm.Set(channel, uint32(float64(pwm.Top())*float64(pulse)/float64(period)))
	}
}