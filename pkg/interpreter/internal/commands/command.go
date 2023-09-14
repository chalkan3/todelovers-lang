package commands

import (
	control "mary_guica/pkg/tvm/pkg/control_plane/requester"
)

type Command interface {
	Execute(instruction byte, threadID int, args ...interface{})
}

type base struct {
	requester control.FlightAttendant
}

func (b *base) Request(fn interface{}) { b.requester.Request(fn) }
