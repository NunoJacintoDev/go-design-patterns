package builder

type ManufacturingDirector struct {
	b BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.b.SetSeats()
	f.b.SetWheels()
	f.b.SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.b = b
}
