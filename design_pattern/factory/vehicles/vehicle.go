package vehicles

import "fmt"

type Car struct{}

func (car Car) Drive() {
	fmt.Println("Driving a car on the road...")
}

type Bicycle struct{}

func (b Bicycle) Drive() {
	fmt.Println("Riving a bicycle on the road...")
}

type Truck struct{}

func (t Truck) Drive() {
	fmt.Println("Driving a truck on the road...")
}

// VehicleFactory is the Factory function to create vehicles
func VehicleFactory(vehicleType string) Vehicle {
	switch vehicleType {
	case "car":
		return Car{}
	case "bicycle":
		return Bicycle{}
	case "truck":
		return Truck{}
	default:
		return nil
	}
}
