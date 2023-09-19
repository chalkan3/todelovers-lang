package metrics

import (
	"fmt"
	"mary_guica/pkg/tvm/pkg/events"
	"runtime"
	"time"
)

type GorotineManager interface {
	Count()
}

type gorotineManager struct {
}

func NewGorotineManager() GorotineManager { return new(gorotineManager) }
func (m *gorotineManager) Count() {
	for {
		events.GetEventController().Notify(&events.Notifier{
			Handler: "NOTIFY",
			Event: &events.Event{
				Name:        "COUNT_GOROTINES",
				Description: fmt.Sprintf("Total of gorotines running: {%d} ", runtime.NumGoroutine()),
				Data:        runtime.NumGoroutine(),
			},
		})
		time.Sleep(5 * time.Second)
	}
}
