package main

import (
	"time"

	spinners "github.com/tom-draper/go-spinners"
)

func main() {
	SpinnerNames := []string{"dots1"}
	s := spinners.Spinner("")
	s.Start()
	for _, name := range SpinnerNames {
		s.SetPrefix(name + ":")
		s.SetSpinner(name)
		time.Sleep(time.Second * 2)
	}
	s.Stop()
}
