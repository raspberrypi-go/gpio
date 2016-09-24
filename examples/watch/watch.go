package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/raspberrypi-go/gpio"
)

func main() {
	// // set GPIO22 to input mode
	pin, err := gpio.OpenPin(gpio.GPIO22, gpio.ModeInput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}

	power, err := gpio.OpenPin(gpio.GPIO15, gpio.ModeOutput)

	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}

	// clean up on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Println("Closing pin and terminating program.")
			power.Set()
			pin.Close()
			os.Exit(0)
		}
	}()

	err = pin.BeginWatch(gpio.EdgeFalling, func() {
		fmt.Printf("Callback for %d triggered!\n\n", gpio.GPIO22)
	})
	if err != nil {
		fmt.Printf("Unable to watch pin: %s\n", err.Error())
		os.Exit(1)
	}

	// fmt.Println("Now watching pin 22 on a falling edge.")

	// set GPIO14 to output mode
	ring := func(step time.Duration) {
		power.Set()
		time.Sleep(step)
		power.Clear()
		time.Sleep(step)
	}

	step := 100
	for {
		// for i := 0; i < 1000; i++ {
		ring(time.Second)
		// }
		// if step < 10000 {
		// 	step += 100
		// } else {
		// 	step = 100
		// }
		fmt.Println(step)
	}

}
