package main

import "fmt"

/*
la struttura snake è formata da una coda di N nodi
ciascuno con delle coordinate
*/
type board struct {
	xy [20][20]string
}

type node struct {
	x, y int
	next *node
}
type snake struct {
	first *node
	last  *node
}

// global variables
var b board

// init the snake with length of 2, centered.
func init_snake(s snake) {

}

func draw(s snake) {

}

func upadate_board(board *board) {

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

	b.xy[0][0] = "*"

	// goroutine che aggiorna ogni delta time
	// la posizione dello snake nella mappa

	// goroutine che attende gli input dell'utente
}
