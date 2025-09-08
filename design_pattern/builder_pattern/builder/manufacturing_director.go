package builder

type ManufacturingDirector struct {
	builder BuilderProcess
}

func (m *ManufacturingDirector) SetBuilder(builder BuilderProcess) {
	m.builder = builder
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}
