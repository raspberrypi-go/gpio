package rpi

import (
	"github.com/raspberrypi-go/gpio"
)

// assert that rpi.pin implements gpio.Pin
var _ gpio.Pin = new(pin)
