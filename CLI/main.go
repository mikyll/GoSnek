package main

import "fmt"

const BL = 10 // Board Length
const BH = 10 // Board Height

/*
la struttura snake è formata da una coda di N nodi
ciascuno con delle coordinate
*/
type board struct {
	xy [10][10]string
}

type node struct {
	x, y int
	next *node
}
type snake struct {
	heading int
	first   *node
}

// global variables
var b board
var s snake

// init the snake with length of 2, centered.
func init_snake() {
	n1 := node{x: 10, y: 10, next: nil}
	s.first = &n1

}

func draw(s snake) {

}

func upadate_board() {
	// nb: ogni nodo si sposta nella posizione del
	// precedente, tranne il primo, che va in
	// direzione di heading

	new_x := s.first.x
	new_y := s.first.y

	b.xy[new_x][new_y] = "x"
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

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 || i == 9 || j == 0 || j == 9 {
				b.xy[i][j] = "*"
			}
		}
	}

	init_snake()

	upadate_board()

	// goroutine che aggiorna ogni delta time
	// la posizione dello snake nella mappa

	// goroutine che attende gli input dell'utente
}
