package metrics

import (
	"log"
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
		log.Printf("[VM] Total current gorotine: %d\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}
