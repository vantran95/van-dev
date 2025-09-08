package main

import (
	"fmt"

	"github/van-dev/design_pattern/builder_pattern/builder"
)

func main() {
	manufacturingVehicle := builder.ManufacturingDirector{}
	bicycleBuilder := &builder.BicycleBuilder{}

	manufacturingVehicle.SetBuilder(bicycleBuilder)
	manufacturingVehicle.Construct()

	bicycle := bicycleBuilder.GetVehicle()
	fmt.Printf("Vehicle is %s has %d wheels, %d seats.", bicycle.Structure, bicycle.Wheels, bicycle.Seats)
}
