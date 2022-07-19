package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type spinner struct {
	mu    sync.Mutex
	chars []string
	text  string
	halt  chan struct{}
	delay time.Duration
}

func chars(name string) []string {
	var chars []string
	switch name {
	case "dots1":
		chars = strings.Split("⠋⠙⠹⠼⠴⠦⠧⠇⠏", "")
	default:
		chars = strings.Split("|/-\\", "")
	}
	return chars
}

func spinnerText(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func Spinner(name string, args ...string) *spinner {
	text := spinnerText(args)
	chars := chars(name)
	return &spinner{
		chars: chars,
		text:  text,
		halt:  make(chan struct{}, 1),
		delay: time.Millisecond * 100,
	}
}

func (s *spinner) Start() {
	go s.animate()
}

func (s *spinner) animate() {
	for {
		for i := 0; i < len(s.chars); i++ {
			select {
			case <-s.halt:
				return
			default:
				s.mu.Lock()
				fmt.Print("\033[?25l") // Hides cursor
				fmt.Print("\r\033[K")  // Erases line
				fmt.Print(s.chars[i])  // Prints char
				fmt.Printf(" %s", s.text)
				delay := s.delay
				s.mu.Unlock()
				time.Sleep(delay)
			}
		}
	}
}

func (s *spinner) Stop() {
	s.mu.Lock()
	fmt.Println()
	s.halt <- struct{}{}
	s.mu.Unlock()
}

func main() {
	s := Spinner("dots1", " Loading...")
	s.Start()
	time.Sleep(time.Second * 10)
	s.Stop()
}
