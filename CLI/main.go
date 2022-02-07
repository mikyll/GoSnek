package main

import "fmt"

const BL = 10 // Board Length
const BH = 10 // Board Height

/*
la struttura snake è formata da una coda di N nodi
ciascuno con delle coordinate
*/
type board struct {
	xy [BL][BH]string
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
var b board
var s snake

// init the snake with length of 2, centered.
func init_snake() {
	n2 := node{x: BL/2 + 1, y: BH / 2, next: nil}
	n1 := node{x: BL / 2, y: BH / 2, next: &n2}
	s.first = &n1

}

func draw() {
	fmt.Printf("---Snake---\n")
	for i := 0; i < BL; i++ {
		for j := 0; j < BH; j++ {
			fmt.Printf("%s", b.xy[i][j])
			if i == BL {
				fmt.Printf("\n")
			}
		}
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
	fmt.Printf("ciao")

	// init board
	for i := 0; i < BL; i++ {
		for j := 0; j < BH; j++ {
			if i == 0 || i == 9 || j == 0 || j == 9 {
				b.xy[i][j] = "*"
			} else {
				b.xy[i][j] = " "
			}
		}
	}

	init_snake()

	draw()

	//upadate_board()

	// goroutine che aggiorna ogni delta time
	// la posizione dello snake nella mappa

	// goroutine che attende gli input dell'utente
}
