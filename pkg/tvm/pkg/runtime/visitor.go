package runtime

type Visitor interface {
	Visit(f FlightAttendant, fn func(fa FlightAttendant))
}

type isRunningVisitor struct{}

func (ir *isRunningVisitor) Visit(f FlightAttendant, fn func(fa FlightAttendant)) {
	if f.Running() {
		return
	}

	go fn(f)
	f.SetRunning()
}

func NewIsRunningVisitor() Visitor {
	return &isRunningVisitor{}
}
