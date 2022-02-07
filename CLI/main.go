package main

import "fmt"

/*
la struttura snake è formata da una coda di N nodi
ciascuno con delle coordinate
*/

type node struct {
	x, y int
	next *node
}
type snake struct {
	first *node
	last  *node
}

// init the snake with length of 2, centered.
func init_snake(s snake) {

}

func draw_snake(s snake) {

}

func board_updater() {

}

func main() {
	fmt.Printf("ciao")
}
