package main

import "fmt"

type mover interface {
	move()
}

type animal struct{}

func (a animal) move() {
	fmt.Println("I walk")
}

type bird struct{}

func (b bird) fly() {
	fmt.Println("I fly")
}

type birdAdapter struct {
	b *bird
}

func (ba birdAdapter) move() {
	ba.b.fly()
}
