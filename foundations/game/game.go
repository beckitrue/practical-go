package main

import (
	"fmt"
)

func main() {
	i1 := Item{10, 20}
	fmt.Println(i1)
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
		Keys: []Key{}, // empty slice - players begin with no keys
	}
	fmt.Printf("p1 %#v\n", p1)
	fmt.Printf("p1.Item.X %d\n", p1.Item.X)
	fmt.Printf("p1.Keys %v\n", p1.Keys)

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

	k := Jade
	fmt.Println("k:", k)
	fmt.Println("key:", Key(4))

	p1.FoundKey(Crystal)
	fmt.Println(p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Copper)
	fmt.Println(p1.Keys)
}

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal use
)

type Key byte

// rule of thumb: accept interfaces, return types
func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
}

// i is called a receiver
// Move is called a method of the Item type
// if you want to modify the struct, you need to pass a pointer
func (i *Item) Move(x, y int) {
	i.X += x
	i.Y += y
}

// add new key to the slice of keys if it's not already there
// and if it is a valid key
func (p *Player) FoundKey(k Key) error {
	if k >= invalidKey {
		return fmt.Errorf("Invalid key")
	}
	if p.Keys == nil {
		println("Keys is nil")
		p.Keys = []Key{k}
	}
	if !containsKey(k, p.Keys) {
		p.Keys = append(p.Keys, k)

	}
	return nil
}

func containsKey(k Key, keys []Key) bool {
	for _, key := range keys {
		if key == k {
			fmt.Printf("Key %s already found k %s\n",
				key, k)
			return true
		}
	}
	return false
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

type Item struct {
	X int
	Y int
}

const (
	maxX = 1000
	maxY = 600
)

type Player struct {
	Name string
	Item // embed Item
	Keys []Key
}
