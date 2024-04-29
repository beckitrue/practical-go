package main

import (
	"fmt"
)

func main() {
	i1 := Item{10, 20}
	fmt.Println(i1)
	//fmt.Println(err)
	fmt.Printf("i1 %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Println(i2)

	i3 := Item{
		Y: 10,
	}
	fmt.Println(i3)

	fmt.Println(NewItem(10, -20))

	p1 := Player{
		Name: "John",
		Item: Item{500, 300},
	}
	fmt.Printf("p1 %#v\n", p1)
	fmt.Printf("p1.Item.X %d\n", p1.Item.X)

	p1.Move(400, 600)
	fmt.Printf("p1 (move): %#v\n", p1)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}
}

// rule of thumb: accept interfaces, return types
func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
}

type Player struct {
	Name string
	Item // embed Item
}

// i is called a receiver
// if you want to modify the struct, you need to pass a pointer
func (i *Item) Move(x, y int) {
	i.X += x
	i.Y += y
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || y < 0 || x > maxX || y > maxY {
		return nil, fmt.Errorf("Invalid coordinates. X min %d, Y min %d, X max %d, Y max %d x %d, y %d", 0, 0, maxX, maxY, x, y)
	}
	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)

type Item struct {
	X int
	Y int
}
