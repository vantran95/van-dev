package main

import "github/van-dev/design_pattern/factory/vehicles"

func main() {
	// Create a car through the factor
	car := vehicles.VehicleFactory("car")
	car.Drive()

	// Create a bicycle through the factory
	bicycle := vehicles.VehicleFactory("bicycle")
	bicycle.Drive()

	// Create a truck through the factory
	truck := vehicles.VehicleFactory("truck")
	truck.Drive()
}
