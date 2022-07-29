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

// Many spinners are taken from cli-spinners
// https://github.com/sindresorhus/cli-spinners
func chars(name string) []string {
	var chars []string
	switch name {
	case "dots1", "dots":
		chars = strings.Split("⠋⠙⠹⠼⠴⠦⠧⠇⠏", "")
	case "dots2":
		chars = strings.Split("⣾⣽⣻⢿⡿⣟⣯⣷", "")
	case "dots3":
		chars = strings.Split("⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓", "")
	case "dots4":
		chars = strings.Split("⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆", "")
	case "dots5":
		chars = strings.Split("⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋", "")
	case "dots6":
		chars = strings.Split("⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁", "")
	case "dots7":
		chars = strings.Split("⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈", "")
	case "dots8":
		chars = strings.Split("⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈", "")
	case "dots9":
		chars = strings.Split("⢹⢺⢼⣸⣇⡧⡗⡏", "")
	case "dots10":
		chars = strings.Split("⢄⢂⢁⡁⡈⡐⡠", "")
	case "dots11":
		chars = strings.Split("⠁⠂⠄⡀⢀⠠⠐⠈", "")
	case "dots12":
		chars = []string{"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀",
			"⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩",
			"⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐",
			"⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩",
			"⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐",
			"⠀⠠", "⠀⢀", "⠀⡀"}
	case "line", "lines":
		chars = strings.Split("|/-\\", "")
	case "hline", "hlines", "line2", "lines2":
		chars = strings.Split("⠂-–—–-", "")
	case "arc":
		chars = strings.Split("◜◠◝◞◡◟", "")
	case "circle":
		chars = strings.Split("◡⊙◠", "")
	case "triangle":
		chars = strings.Split("◢◣◤◥", "")
	case "pipe":
		chars = strings.Split("┤┘┴└├┌┬┐", "")
	case "elipses":
		chars = []string{".  ", ".. ", "...", "   "}
	case "elipses2":
		chars = []string{".  ", ".. ", "...", " ..", "  .", "   "}
	case "elipses3":
		chars = []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}
	case "balloon":
		chars = strings.Split(" .oO@*", "")
	case "balloon2":
		chars = strings.Split(".oO°Oo.", "")
	case "noise":
		chars = strings.Split("▓▒░", "")
	case "bounce":
		chars = strings.Split("⠁⠂⠄⠂", "")
	case "bouncingball":
		chars = []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)",
			"(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"}
	case "bouncingball2":
		chars = []string{" ●    ", "  ●   ", "   ●  ", "    ● ", "     ●",
			"    ● ", "   ●  ", "  ●   ", " ●    ", "●     "}
	case "bouncingbar":
		chars = []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]",
			"[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]",
			"[==  ]", "[=   ]"}
	case "boxbounce":
		chars = strings.Split("▖▘▝▗", "")
	case "boxbounce2":
		chars = strings.Split("▌▀▐▄", "")
	case "circlequarters":
		chars = strings.Split("◴◷◶◵", "")
	case "circlehalves":
		chars = strings.Split("◐◓◑◒", "")
	case "squarecorners":
		chars = strings.Split("◰◳◲◱", "")
	case "arrow", "arrows":
		chars = strings.Split("←↖↑↗→↘↓↙", "")
	case "arrow2", "arrows2":
		chars = []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"}
	case "star", "stars":
		chars = strings.Split("✶✸✹✺✹✷", "")
	case "star2", "stars2":
		chars = strings.Split("+x*", "")
	case "flip":
		chars = strings.Split("___-``'´-___", "")
	case "hamburger":
		chars = strings.Split("☱☲☴", "")
	case "vgrow":
		chars = strings.Split("▁▃▄▅▆▇▆▅▄▃", "")
	case "hgrow":
		chars = strings.Split("▏▎▍▌▋▊▉▊▋▌▍▎", "")
	case "cross":
		chars = strings.Split("╫╪", "")
	case "layer":
		chars = strings.Split("-=≡", "")
	case "granade":
		chars = []string{"،   ", "′   ", " ´ ", " ‾ ", "  ⸌", "  ⸊", "  |", "  ⁎",
			"  ⁕", " ෴ ", "  ⁓", "   ", "   ", "   "}
	case "toggle1", "toggle":
		chars = strings.Split("⊶⊷", "")
	case "toggle2":
		chars = strings.Split("▫▪", "")
	case "toggle3":
		chars = strings.Split("□■", "")
	case "toggle4":
		chars = strings.Split("■□▪▫", "")
	case "toggle5":
		chars = strings.Split("▮▯", "")
	case "toggle6":
		chars = strings.Split("⦾⦿", "")
	case "toggle7":
		chars = strings.Split("◍◌", "")
	case "toggle8":
		chars = strings.Split("㊂㊀㊁", "")
	case "toggle9":
		chars = strings.Split("⧇⧆", "")
	case "toggle10":
		chars = strings.Split("=*-", "")
	case "bell":
		chars = strings.Split("☗☖", "")
	case "dqpb":
		chars = strings.Split("dqpb", "")
	case "clock":
		chars = strings.Split("🕛🕐🕑🕒🕓🕔🕕🕖🕗🕘🕙🕚", "")
	case "earth":
		chars = strings.Split("🌍🌎🌏", "")
	case "moon":
		chars = strings.Split("🌑🌒🌓🌔🌕🌖🌗🌘", "")
	case "runner":
		chars = strings.Split("🚶🏃", "")
	case "pong":
		chars = []string{"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌",
			"▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌",
			"▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌",
			"▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌",
			"▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌",
			"▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌",
			"▐⠠       ▌"}
	case "shark":
		chars = []string{"▐|\\____________▌", "▐_|\\___________▌", "▐__|\\__________▌",
			"▐___|\\_________▌", "▐____|\\________▌", "▐_____|\\_______▌",
			"▐______|\\______▌", "▐_______|\\_____▌", "▐________|\\____▌",
			"▐_________|\\___▌", "▐__________|\\__▌", "▐___________|\\_▌",
			"▐____________|\\▌", "▐____________/|▌", "▐___________/|_▌",
			"▐__________/|__▌", "▐_________/|___▌", "▐________/|____▌",
			"▐_______/|_____▌", "▐______/|______▌", "▐_____/|_______▌",
			"▐____/|________▌", "▐___/|_________▌", "▐__/|__________▌",
			"▐_/|___________▌", "▐/|____________▌"}
	default:
		chars = strings.Split("|/-\\", "") // Default to lines
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
	fmt.Print("\r\033[K")
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
				fmt.Printf("%s%s%s", s.prefix, s.chars[i], s.postfix) // Print
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

func (s *spinner) SetDelay(delay time.Duration) {
	s.mu.Lock()
	s.delay = delay
	s.mu.Unlock()
}

func (s *spinner) SetPrefix(prefix string) {
	s.mu.Lock()
	s.prefix = prefix + " "
	s.mu.Unlock()
}

func (s *spinner) SetPostfix(postfix string) {
	s.mu.Lock()
	s.postfix = " " + postfix
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
