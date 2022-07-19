package main

import (
	"time"

	spinners "github.com/tom-draper/go-spinners"
)

var SpinnerNames = []string{"dots1", "dots2", "dots3", "dots4", "dots5", "dots6",
	"dots7", "dots8", "dots9", "dots10", "dots11", "dots12", "line", "hline", "arc",
	"circle", "triangle", "pipe", "elipses", "elipses2", "elipses3", "balloon",
	"noise", "bounce", "bouncingball", "bouncingbar", "boxbounce", "boxbounce2",
	"circlequarters", "circlehalves", "squarecorners", "arrow", "arrow2", "star",
	"star2", "flip", "hamburger", "vgrow", "hgrow", "cross", "layer", "granade",
	"toggle1", "toggle2", "toggle3", "toggle4", "toggle5", "toggle6", "toggle7",
	"toggle8", "toggle9", "toggle10", "toggle11", "dqpb", "clock", "earth", "moon",
	"runner", "pong", "shark"}

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
