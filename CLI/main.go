package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
	"golang.org/x/term"
)

// CONSTANTS ========================================================
const UP = "w"
const DOWN = "s"
const LEFT = "a"
const RIGHT = "d"
const PAUSE = "p"
const ESC = "q"
const F_POINTS = 10
const S_POINTS = 100

// color scheme 0 (default)
const BORDER0 = "*"
const BLANK0 = " "
const HEAD0 = "o"
const BODY0 = "x"
const F0 = "F"
const S0 = "S"

// color schemes 1
const BORDER1 = "\033[37;47m \033[0m"
const BLANK1 = "\033[30;40m \033[0m"
const HEAD1 = "\033[97;107m \033[0m"
const BODY1 = "\033[37;47m \033[0m"
const F1 = "F"
const S1 = "S"

// color scheme 2
const BORDER2 = "\033[37;47m \033[0m"
const BLANK2 = "\033[30;40m \033[0m"
const HEAD2 = "\033[32;42m \033[0m"
const BODY2 = "\033[92;102m \033[0m"
const F2 = "\033[30;41mF\033[0m"
const S2 = "\033[30;43mS\033[0m"

// STRUCTURES =======================================================
type board struct {
	xy [][]string
}

type node struct {
	x, y int
	next *node
}
type snake struct {
	hx, hy int // snake direction
	first  *node
}

type fruit struct {
	x, y  int
	value string
}

// GLOBAL VARIABLES =================================================
var OS = "" // Operating System
var BL = 0  // Board Length
var BH = 0  // Board Height

var COLOR_SCHEME = 2

// board elements
var BORDER = "*"
var BLANK = " "
var HEAD = "o"
var BODY = "x"
var F = "F"
var S = "S"

var game_over = false
var tot_points = 0

var b *board = new(board)
var s *snake = new(snake)
var f *fruit = new(fruit)

var input_channel = make(chan string, 5)

// FUNCTIONS ========================================================
func init_board() {
	b.xy = make([][]string, BL)
	for i := range b.xy {
		b.xy[i] = make([]string, BH)
	}

	for i := 0; i < BL; i++ {
		for j := 0; j < BH; j++ {
			if i == 0 || i == BL-1 || j == 0 || j == BH-1 {
				b.xy[i][j] = BORDER
			} else {
				b.xy[i][j] = BLANK
			}
		}
	}
}

// init the snake with length of 4, in the center of the screen
func init_snake() {
	s.hx = -1
	s.hy = 0
	n4 := node{x: BL/2 + 1, y: BH / 2, next: nil}
	n3 := node{x: BL / 2, y: BH / 2, next: &n4}
	n2 := node{x: BL/2 - 1, y: BH / 2, next: &n3}
	n1 := node{x: BL/2 - 2, y: BH / 2, next: &n2}
	s.first = &n1
}

// spawn fruit in a random position inside the board
func spawn_fruit() {
	f.x = rand.Intn(BL-2) + 1
	f.y = rand.Intn(BH-2) + 1

	// spawn algorithm (to prevent the fruit from spawning inside the snake)
	/*
		We get a random int (starting point, if it's one the snake, we start iterating until we get an empty position)
	*/
	if b.xy[f.x][f.y] == HEAD || b.xy[f.x][f.y] == BODY {
		if f.x == 1 || f.x == BL-1 {
			f.x = BL / 2
		}
		if f.y == 1 || f.y == BH-1 {
			f.y = BH / 2
		}
		found := false
		j := f.y
		i := f.x
		for f.y += 1; f.y != j; f.y++ {
			for f.x += 1; f.x != i; f.x++ {
				if b.xy[f.x][f.y] != HEAD && b.xy[f.x][f.y] != BODY && b.xy[f.x][f.y] != BORDER {
					found = true
					break
				}
				if f.x == BL-1 {
					f.x = 1
				}
			}
			if found {
				break
			}
			if f.y == BH-1 {
				f.y = 1
			}
		}
	}

	if rand.Intn(100) < 10 {
		f.value = S
	} else {
		f.value = F
	}
	b.xy[f.x][f.y] = f.value
}

// add a snake node on the head
func add_node(x, y int) {
	n := node{x: x, y: y, next: s.first}
	s.first = &n
}

func update_snake_position() {
	// checks for collision with snake
	if b.xy[s.first.x+s.hx][s.first.y+s.hy] == BODY {
		game_over = true
		return
	}
	add_node(s.first.x+s.hx, s.first.y+s.hy)

	// checks for collision with borders
	if s.first.x == 0 || s.first.x == BL-1 || s.first.y == 0 || s.first.y == BH-1 {
		game_over = true
		return
	}

	prev_node := s.first
	node := s.first.next
	for {
		if node.next != nil {
			prev_node = node
			node = node.next
		} else {
			b.xy[node.x][node.y] = BLANK
			prev_node.next = nil
			break
		}
	}
}

// updates the snake inside the board
func update_board() {
	var node *node

	node = s.first.next
	b.xy[s.first.x][s.first.y] = HEAD

	for {
		b.xy[node.x][node.y] = BODY
		if node.next != nil {
			node = node.next
		} else {
			break
		}
	}
}

// check if the snake has collected a fruit
func collect_fruit() {
	if s.first.x == f.x && s.first.y == f.y {
		if f.value == F {
			tot_points += F_POINTS
		} else if f.value == S {
			tot_points += S_POINTS
		} else {
			// Error
			return
		}
		add_node(f.x, f.y)
		spawn_fruit()
	}
}

func draw() {
	fmt.Printf("\033[1;0H")
	for y := 0; y < BH; y++ {
		for x := 0; x < BL; x++ {
			fmt.Printf("%s", b.xy[x][y])
		}
		fmt.Printf("\n")
	}
}

/*
 SSSS  N   N   AAAA  K   K  EEEEE
S      NN  N  A   A  K  K   E
 SSS   N N N  A   A  KKK    EEE
    S  N  NN  AAAAA  K  K   E
SSSS   N   N  A   A  K   K  EEEEE
*/

func print_controls() {
	fmt.Print("\033[H\033[2J")
	for j := 0; j < BH; j++ {
		for i := 0; i < BL; i++ {
			switch {
			case i == 0 || i == BL-1 || j == 0 || j == BH-1:
				fmt.Print(BORDER)

			case j == 2 && i == BL/2-16:
				fmt.Printf(" \033[37;47mSSSS\033[0m  \033[37;47mN\033[0m   \033[37;47mN\033[0m   \033[37;47mAAAA\033[0m  \033[37;47mK\033[0m   \033[37;47mK\033[0m  \033[37;47mEEEEE\033[0m")
				i += 32
			case j == 3 && i == BL/2-16:
				fmt.Printf("\033[37;47mS\033[0m      \033[37;47mNN\033[0m  \033[37;47mN\033[0m  \033[37;47mA\033[0m   \033[37;47mA\033[0m  \033[37;47mK\033[0m  \033[37;47mK\033[0m   \033[37;47mE\033[0m    ")
				i += 32
			case j == 4 && i == BL/2-16:
				fmt.Printf(" \033[37;47mSSS\033[0m   \033[37;47mN\033[0m \033[37;47mN\033[0m \033[37;47mN\033[0m  \033[37;47mA\033[0m   \033[37;47mA\033[0m  \033[37;47mKKK\033[0m    \033[37;47mEEE\033[0m  ")
				i += 32
			case j == 5 && i == BL/2-16:
				fmt.Printf("    \033[37;47mS\033[0m  \033[37;47mN\033[0m  \033[37;47mNN\033[0m  \033[37;47mAAAAA\033[0m  \033[37;47mK\033[0m  \033[37;47mK\033[0m   \033[37;47mE\033[0m    ")
				i += 32
			case j == 6 && i == BL/2-16:
				fmt.Printf("\033[37;47mSSSS\033[0m   \033[37;47mN\033[0m   \033[37;47mN\033[0m  \033[37;47mA\033[0m   \033[37;47mA\033[0m  \033[37;47mK\033[0m   \033[37;47mK\033[0m  \033[37;47mEEEEE\033[0m")
				i += 32

			case i == BL/2-2 && j == BH/2:
				fmt.Print(HEAD)
				k := 0
				for node := s.first; node.next != nil; k++ {
					fmt.Print(BODY)
					node = node.next
				}
				i += k

			case j == BH/2+2 && i == BL/2-4:
				fmt.Printf("CONTROLS")
				i += 7
			case j == BH/2+4 && i == BL/2-10:
				fmt.Printf("\U00002191 w = up")
				i += 7
			case j == BH/2+5 && i == BL/2-10:
				fmt.Printf("\U00002193 s = down")
				i += 9
			case j == BH/2+6 && i == BL/2-10:
				fmt.Printf("\U00002190 a = left")
				i += 9
			case j == BH/2+7 && i == BL/2-10:
				fmt.Printf("\U00002192 d = right")
				i += 10
			case j == BH/2+5 && i == BL/2+5:
				fmt.Printf("p = pause")
				i += 8
			case j == BH/2+6 && i == BL/2+5:
				fmt.Printf("q = quit")
				i += 7
			default:
				fmt.Print(BLANK)
			}
		}
		fmt.Printf("\n")
	}
}

func print_game_over() {
	fmt.Print("\033[H\033[2J")
	for j := 0; j < BH; j++ {
		for i := 0; i < BL; i++ {
			switch {
			case i == 0 || i == BL-1 || j == 0 || j == BH-1:
				fmt.Print(BORDER)
			case j == BH/2-1 && i == BL/2-5:
				fmt.Printf("GAME OVER")
				i += 8
			case j == BH/2+1 && i == BL/2-(8+(len(strconv.Itoa(tot_points))/2)):
				fmt.Printf("Total Points: %d", tot_points)
				i += 13 + len(strconv.Itoa(tot_points))
			default:
				fmt.Print(BLANK)
			}
		}
		if j < BH-1 {
			fmt.Printf("\n")
		}
	}
}

func show_points() {
	fmt.Printf("Points: %d", tot_points)
}

// GOROUTINES =======================================================
func game() {

	for !game_over {
		update_snake_position()
		update_board()
		collect_fruit()
		draw()
		show_points()
		// check if there are inputs
		select {
		case x := <-input_channel:
			switch x {
			case UP:
				if s.hy != 1 {
					s.hx = 0
					s.hy = -1
				}
			case DOWN:
				if s.hy != -1 {
					s.hx = 0
					s.hy = +1
				}
			case LEFT:
				if s.hx != 1 {
					s.hx = -1
					s.hy = 0
				}
			case RIGHT:
				if s.hx != -1 {
					s.hx = +1
					s.hy = 0
				}
			case PAUSE:
				for i := -4; i < 2; i++ {
					b.xy[BL/2+i][BH/2-1] = " "
				}
				b.xy[BL/2-4][BH/2] = " "
				b.xy[BL/2-3][BH/2] = "P"
				b.xy[BL/2-2][BH/2] = "A"
				b.xy[BL/2-1][BH/2] = "U"
				b.xy[BL/2][BH/2] = "S"
				b.xy[BL/2+1][BH/2] = "E"
				b.xy[BL/2+2][BH/2] = " "
				for i := -4; i < 2; i++ {
					b.xy[BL/2+i][BH/2+1] = " "
				}
				draw()
				x = <-input_channel
				if x == ESC {
					return
				}
				b.xy[BL/2-3][BH/2] = " "
				b.xy[BL/2-2][BH/2] = " "
				b.xy[BL/2-1][BH/2] = " "
				b.xy[BL/2][BH/2] = " "
				b.xy[BL/2+1][BH/2] = " "
			case ESC:
				return
			default:
				fmt.Printf("[INPUT] Input %s not valid. Press 'q' to quit\n", x)
			}
		default:
			continue
		}
	}
}

func input_sampler() {
	// switch stdin into 'raw' mode
	switch {
	case OS == "windows":
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)
	case OS == "darwin" || OS == "linux":
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run() // disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()              // do not display entered characters on the screen
		defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	default:
		fmt.Printf("%s.\n", OS)
		return
	}
	ch := make([]byte, 1)
	var err error
	for {
		// read byte
		_, err = os.Stdin.Read(ch)
		if err != nil {
			fmt.Println(err)
			return
		}

		// send on channel
		input_channel <- string(ch[0])

	}
}

// MAIN =============================================================
func main() {
	// Set color scheme
	BORDER = BORDER2
	BLANK = BLANK2
	HEAD = HEAD2
	BODY = BODY2
	F = F2
	S = S2

	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	// Init termbox
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	// Get window size
	w, h := termbox.Size()
	termbox.Close()
	fmt.Println(w, h) // test
	BL = w
	BH = h - 1

	// Check OS
	OS = runtime.GOOS

	// Switch stdin into 'raw' mode
	switch {
	case OS == "windows":
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)
	case OS == "darwin" || OS == "linux":
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run() // disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()              // do not display entered characters on the screen
		defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	default:
		fmt.Printf("%s.\n", OS)
		return
	}

	// Init board
	init_board()
	init_snake()
	spawn_fruit()

	// Show controls
	print_controls()
	ch := make([]byte, 1)
	_, err := os.Stdin.Read(ch)
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(ch[0]) == "q" {
		return
	}

	// Start game
	go input_sampler()
	game()

	// Game Over
	print_game_over()
}
