package main

import "fmt"

// IDraw IDraw
type IDraw interface {
	Draw() string
}

// Square 正方形
type Square struct{}

// Draw
func (s Square) Draw() string {
	return "this is a square"
}

// ColorSquare 有颜色的正方形
type ColorSquare struct {
	square IDraw
	color  string
}

// NewColorSquare
func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{color: color, square: square}
}

// Draw
func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}
func main() {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	fmt.Println(got)
}
