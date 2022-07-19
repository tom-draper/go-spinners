package spinner

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type spinner struct {
	mu      sync.Mutex
	chars   []string
	postfix string
	prefix  string
	halt    chan struct{}
	delay   time.Duration
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

func Spinner(name string) *spinner {
	chars := chars(name)
	return &spinner{
		chars:   chars,
		postfix: "",
		prefix:  "",
		halt:    make(chan struct{}, 1),
		delay:   time.Millisecond * 100,
	}
}

func (s *spinner) Start() {
	go s.animate()
}

func (s *spinner) eraseLine() {
	fmt.Print("\r\033[K") // Erases line
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
				s.eraseLine()
				fmt.Printf("%s %s %s", s.prefix, s.chars[i], s.postfix) // Print
				delay := s.delay
				s.mu.Unlock()
				time.Sleep(delay)
			}
		}
	}
}

func (s *spinner) Stop() {
	s.mu.Lock()
	s.eraseLine()
	s.halt <- struct{}{}
	s.mu.Unlock()
}

func (s *spinner) SetPrefix(prefix string) {
	s.mu.Lock()
	s.prefix = prefix
	s.mu.Unlock()
}

func (s *spinner) SetPostfix(postfix string) {
	s.mu.Lock()
	s.postfix = postfix
	s.mu.Unlock()
}

func (s *spinner) SetSpinner(name string) {
	s.mu.Lock()
	s.chars = chars(name)
	s.mu.Unlock()
}

func main() {
	s := Spinner("dots1")
	s.Start()
	time.Sleep(time.Second * 10)
	s.Stop()
}
