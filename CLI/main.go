package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"golang.org/x/term"
)

const W = "w"
const S = "s"
const A = "a"
const D = "d"
const ESC = "q"
const F_POINTS = 10
const S_POINTS = 100

var BL = 120 // Board Length
var BH = 10  // Board Height

var game_over = false
var tot_points = 0

var input_channel = make(chan string, 5)

/*
la struttura snake è formata da una coda di N nodi
ciascuno con delle coordinate
*/
type board struct {
	xy [][]string
}

type node struct {
	x, y int
	next *node
}
type snake struct {
	hx, hy int // the node where the snake is going
	first  *node
}

type fruit struct {
	x, y  int
	value string
}

// global variables
var b *board = new(board)
var s *snake = new(snake)
var f *fruit = new(fruit)

// init the snake with length of 2, centered.
func init_snake() {
	s.hx = -1
	s.hy = 0
	n4 := node{x: BL/2 + 3, y: BH / 2, next: nil}
	n3 := node{x: BL/2 + 2, y: BH / 2, next: &n4}
	n2 := node{x: BL/2 + 1, y: BH / 2, next: &n3}
	n1 := node{x: BL / 2, y: BH / 2, next: &n2}
	s.first = &n1
}

func spawn_fruit() {
	f.x = rand.Intn(BL-2) + 1
	f.y = rand.Intn(BH-2) + 1
	if rand.Intn(100) < 10 {
		f.value = "S"
	} else {
		f.value = "F"
	}
	b.xy[f.x][f.y] = f.value
}

func collect_fruit() {
	if s.first.x == f.x && s.first.y == f.y {
		if f.value == "F" {
			tot_points += F_POINTS
		} else {
			tot_points += S_POINTS
		}
		add_node(f.x, f.y)
		spawn_fruit()
	}
}

func add_node(x, y int) {
	n := node{x: x, y: y, next: s.first}
	s.first = &n
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

func upadate_board() {
	// nb: ogni nodo si sposta nella posizione del
	// precedente, tranne il primo, che va in
	// direzione di heading
	var node *node

	node = s.first.next
	b.xy[s.first.x][s.first.y] = "o"

	for {
		b.xy[node.x][node.y] = "x"
		if node.next != nil {

			node = node.next
		} else {
			break
		}

	}

}

func update_snake_position() {
	// collision with snake
	if b.xy[s.first.x+s.hx][s.first.y+s.hy] == "x" {
		game_over = true
		return
	}
	n := node{x: s.first.x + s.hx, y: s.first.y + s.hy, next: s.first}
	s.first = &n

	// collision with borders
	if n.x == 0 || n.x == BL-1 || n.y == 0 || n.y == BH-1 {
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
			b.xy[node.x][node.y] = " "
			prev_node.next = nil
			break
		}
	}
}

func show_points() {
	fmt.Printf("Points: %d", tot_points)
}

// goroutines
func game() {

	for !game_over {
		update_snake_position()
		upadate_board()
		collect_fruit()
		draw()
		show_points()
		// check if there are inputs
		select {
		case x := <-input_channel:
			switch x {
			case W:
				if s.hy != 1 {
					s.hx = 0
					s.hy = -1
				}
			case S:
				if s.hy != -1 {
					s.hx = 0
					s.hy = +1
				}
			case A:
				if s.hx != 1 {
					s.hx = -1
					s.hy = 0
				}
			case D:
				if s.hx != -1 {
					s.hx = +1
					s.hy = 0
				}
			case ESC:
				return
			default:
				fmt.Printf("[INPUT] Input %s not valid.\n", x)
			}
		default:
			continue
		}
	}
}

func input_sampler() {
	for {
		// read char
		ch := make([]byte, 1)
		_, err := os.Stdin.Read(ch)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(string(ch[0]))

		// send on channel
		input_channel <- string(ch[0])

	}
}
func print_snake() {
	node := s.first
	for {
		fmt.Printf("%v, ", node)
		if node.next != nil {
			node = node.next
		} else {
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	termbox.Close()
	//fmt.Println(w, h) // test
	BL = w
	BH = h - 1

	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// read char
	fmt.Printf("\n---- CONTROLS ----\nw = up\ns = down\na = left\nd = right\n\nq = quit\n\nPress any key to start ...")
	ch := make([]byte, 1)
	_, err = os.Stdin.Read(ch)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(ch[0]))

	// init board
	b.xy = make([][]string, BL)
	for i := range b.xy {
		b.xy[i] = make([]string, BH)
	}

	for i := 0; i < BL; i++ {
		for j := 0; j < BH; j++ {
			if i == 0 || i == BL-1 || j == 0 || j == BH-1 {
				b.xy[i][j] = "*"
			} else {
				b.xy[i][j] = " "
			}
		}
	}

	init_snake()
	//fmt.Printf("[MAIN] Init snake\n")

	spawn_fruit()

	update_snake_position()
	//fmt.Printf("[MAIN] Update snake position\n")

	upadate_board()

	draw()

	//fmt.Printf("[MAIN] Draw board\n")

	//upadate_board()
	go input_sampler()
	game()

	fmt.Print("\033[H\033[2J")
	fmt.Printf("\n\n\nGAME OVER\n\n\nTotal points: %d\n\n\n", tot_points)

	// goroutine che aggiorna ogni delta time
	// la posizione dello snake nella mappa

	// goroutine che attende gli input dell'utente
}
