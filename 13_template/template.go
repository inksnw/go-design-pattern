package main

import "fmt"

type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
	Play()
}

type BaseGame struct {
}

func (b *BaseGame) Initialize() {
	fmt.Println("base Initialize")
}
func (b *BaseGame) StartPlay() {
	fmt.Println("base StartPlay")
}
func (b *BaseGame) EndPlay() {
	fmt.Println("base EndPlay")
}
func (b *BaseGame) Play() {
	b.Initialize()
	b.StartPlay()
	b.EndPlay()
}

type Cricket struct {
	baseGame BaseGame
}

func (c *Cricket) Initialize() {
	fmt.Println("Cricket Game Finished!")
}
func (c *Cricket) StartPlay() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}
func (c *Cricket) EndPlay() {
	c.baseGame.EndPlay()
}
func (c *Cricket) Play() {
	c.Initialize()
	c.StartPlay()
	c.EndPlay()
}

func main() {
	cricket := new(Cricket)
	cricket.Play()
}
