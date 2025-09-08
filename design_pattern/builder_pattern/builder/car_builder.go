package builder

type CarBuilder struct {
	vehicle VehicleProduct
}

func (c *CarBuilder) SetWheels() BuilderProcess {
	c.vehicle.Wheels = 4

	return c
}

func (c *CarBuilder) SetSeats() BuilderProcess {
	c.vehicle.Seats = 4

	return c
}

func (c *CarBuilder) SetStructure() BuilderProcess {
	c.vehicle.Structure = "Car"

	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicle
}
