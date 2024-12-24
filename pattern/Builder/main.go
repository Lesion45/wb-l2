package main

import "fmt"

type House struct {
	Walls   string
	Roof    string
	Windows int
	Doors   int
}

type Builder interface {
	SetWalls()
	SetRoof()
	SetWindows()
	SetDoors()
	GetResult() House
}

type WoodenHouseBuilder struct {
	house House
}

func (b *WoodenHouseBuilder) SetWalls() {
	b.house.Walls = "Wooden walls"
}

func (b *WoodenHouseBuilder) SetRoof() {
	b.house.Roof = "Wooden roof"
}

func (b *WoodenHouseBuilder) SetWindows() {
	b.house.Windows = 4
}

func (b *WoodenHouseBuilder) SetDoors() {
	b.house.Doors = 2
}

func (b *WoodenHouseBuilder) GetResult() House {
	return b.house
}

type StoneHouseBuilder struct {
	house House
}

func (b *StoneHouseBuilder) SetWalls() {
	b.house.Walls = "Stone walls"
}

func (b *StoneHouseBuilder) SetRoof() {
	b.house.Roof = "Stone roof"
}

func (b *StoneHouseBuilder) SetWindows() {
	b.house.Windows = 6
}

func (b *StoneHouseBuilder) SetDoors() {
	b.house.Doors = 3
}

func (b *StoneHouseBuilder) GetResult() House {
	return b.house
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) Construct() {
	d.builder.SetWalls()
	d.builder.SetRoof()
	d.builder.SetWindows()
	d.builder.SetDoors()
}

func main() {
	director := &Director{}

	woodenBuilder := &WoodenHouseBuilder{}
	director.SetBuilder(woodenBuilder)
	director.Construct()
	woodenHouse := woodenBuilder.GetResult()
	fmt.Printf("Wooden house %+v\n", woodenHouse)

	stoneBuilder := &StoneHouseBuilder{}
	director.SetBuilder(stoneBuilder)
	director.Construct()
	stoneHouse := stoneBuilder.GetResult()
	fmt.Printf("Stone house: %+v\n", stoneHouse)
}
