package builder

type BicycleBuilder struct {
	vehicle VehicleProduct
}

func (b *BicycleBuilder) SetWheels() BuilderProcess {
	b.vehicle.Wheels = 2

	return b
}

func (b *BicycleBuilder) SetSeats() BuilderProcess {
	b.vehicle.Seats = 2

	return b
}

func (b *BicycleBuilder) SetStructure() BuilderProcess {
	b.vehicle.Structure = "Bicycle"

	return b
}

func (b *BicycleBuilder) GetVehicle() VehicleProduct {
	return b.vehicle
}
