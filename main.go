package main

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	fmt.Println("Hello, World!")
	area, _ := pterm.DefaultArea.WithCenter().Start()

	for i := 0; i < 5; i++ {
			// Update the content of the area with the current count.
			// The Sprintf function is used to format the string.
			area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

			// Pause for a second to simulate a time-consuming task.
			time.Sleep(time.Second)
	}
}
