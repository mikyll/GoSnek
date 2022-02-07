package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

var BL = 120 // Board Length
var BH = 10  // Board Height

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

// global variables
var b *board = new(board)
var s *snake = new(snake)

// init the snake with length of 2, centered.
func init_snake() {
	s.hx = BL/2 - 1
	s.hy = BH / 2
	n2 := node{x: BL/2 + 1, y: BH / 2, next: nil}
	n1 := node{x: BL / 2, y: BH / 2, next: &n2}
	s.first = &n1

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

}

func update_snake_position() {
	b.xy[s.hx][s.hy] = "x"
	px := s.first.x
	py := s.first.y
	node := s.first.next
	b.xy[px][py] = "x"
	for {
		if node.next != nil {
			px = node.x
			py = node.y
			node = node.next
			b.xy[px][py] = "x"
		} else {
			break
		}
	}
}

// goroutines
func updater() {

}

func input_sampler() {
	for {
		// read char
		ch := 1
		switch ch {
		case 'W':
		case 'S':
		case 'A':
		case 'D':
		default:
			fmt.Printf("[INPUT] Input %d not valid.\n", ch)
		}
	}
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	termbox.Close()
	fmt.Println(w, h)
	BL = w
	BH = h - 1

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

	update_snake_position()
	//fmt.Printf("[MAIN] Update snake position\n")

	draw()
	//fmt.Printf("[MAIN] Draw board\n")

	//upadate_board()

	// goroutine che aggiorna ogni delta time
	// la posizione dello snake nella mappa

	// goroutine che attende gli input dell'utente
}
