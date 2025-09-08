package builder

type BuilderProcess interface {
	SetWheels() BuilderProcess
	SetSeats() BuilderProcess
	SetStructure() BuilderProcess
	GetVehicle() VehicleProduct
}
