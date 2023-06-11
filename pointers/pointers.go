package main

import "fmt"

func zero(x *int) {
	*x = 0
}

type Position struct {
	x float32
	y float32
}

type BadGuy struct {
	name string
	health int
	pos Position
}

func setbadGuyPosition(x, y float32, villain *BadGuy) {
	villain.pos.x = x
	villain.pos.y = y
}

func main() {
	x := 5

	zero(&x)
	// fmt.Println(x)
	var SideVillain BadGuy
	setbadGuyPosition(32, 42, &SideVillain)
	fmt.Println(SideVillain)
}
