package main

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Driving a car.")
}

type Bike struct{}

func (b *Bike) Drive() {
	fmt.Println("Riding a bike.")
}

type VehicleFactory interface {
	CreateVehicle() Vehicle
}

type CarFactory struct{}

func (c *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}

type BikeFactory struct{}

func (b *BikeFactory) CreateVehicle() Vehicle {
	return &Bike{}
}

func main() {
	var vehicleFactory VehicleFactory

	neededVehicle := "Car"

	switch neededVehicle {
	case "Car":
		vehicleFactory = &CarFactory{}
		vehicle := vehicleFactory.CreateVehicle()
		vehicle.Drive()
	case "Bike":
		vehicleFactory = &BikeFactory{}
		vehicle := vehicleFactory.CreateVehicle()
		vehicle.Drive()
	default:
		fmt.Println("Our fabric doesn't support this type of vehicles!")
	}
}
