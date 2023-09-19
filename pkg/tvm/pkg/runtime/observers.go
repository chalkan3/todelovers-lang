package runtime

import (
	"mary_guica/pkg/tvm/pkg/events"
)

type ReloadCrewObserver struct {
	crew Crew
}

func NewReloadCrewObserver(crew Crew) *ReloadCrewObserver {
	return &ReloadCrewObserver{
		crew: crew,
	}
}

func (cl *ReloadCrewObserver) Update(event events.Event) {
	cl.crew.Handler()

}
