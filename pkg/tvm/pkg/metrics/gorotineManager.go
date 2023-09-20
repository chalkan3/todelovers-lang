package metrics

import (
	"fmt"
	"mary_guica/pkg/nando"
	eapi "mary_guica/pkg/tvm/internal/api/events"
	"mary_guica/pkg/tvm/pkg/events"
	"time"

	"runtime"
)

type GorotineManager interface {
	Count()
}

type gorotineManager struct {
}

func NewGorotineManager() GorotineManager { return new(gorotineManager) }
func (m *gorotineManager) Count() {
	c := &nando.Client{}
	for {
		time.Sleep(3 * time.Second)
		c.Do(nando.NewRequest(eapi.Notify.String(), &eapi.NotifyRequest{
			Notifier: &events.Notifier{
				Handler: "NOTIFY",
				Event: &events.Event{
					Name:        "COUNT_GOROTINES",
					Description: fmt.Sprintf("Total of gorotines running: {%d} ", runtime.NumGoroutine()),
					Data:        runtime.NumGoroutine(),
				},
			},
		}))

	}

}
