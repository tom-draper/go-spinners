package main

import (
	"time"

	spinners "github.com/tom-draper/go-spinners"
)

var SpinnerNames = []string{"dots1", "dots2", "dots3", "dots4", "dots5", "dots6",
	"dots7", "dots8", "dots9", "dots10", "dots11", "dots12", "line", "hline", "arc",
	"circle", "triangle", "pipe", "elipses", "elipses2", "elipses3", "balloon", "noise", "bounce", "bouncingball", "bouncingbar", "boxbounce", "boxbounce2", "circlequarters", "runner", "circlehalves", "squarecorners", "arrow", "arrow2", "dqpb", "clock", "earth", "moon", "pong", "shark"}

func main() {
	s := spinners.Spinner("")
	s.Start()
	for _, name := range SpinnerNames {
		s.SetPrefix(name + ":")
		s.SetSpinner(name)
		time.Sleep(time.Second * 2)
	}
	s.Stop()
}
