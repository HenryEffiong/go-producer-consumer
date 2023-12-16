package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// Keep track of Pizza

	// Run forever making pizzas until we quit
	for {
		// Try to make a pizza
	}
}

func main() {
	//  seed the random number generator
	rand.New(rand.NewSource(time.Now().Unix()))

	//  print out a message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("---------------------------------")

	//  create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//  run the producer in the background
	go pizzeria(pizzaJob)

	//  create and run consumer

	//  print out the ending message
}
