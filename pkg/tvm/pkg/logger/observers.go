package logger

import (
	"fmt"
	"io"
	"mary_guica/pkg/tvm/pkg/events"
	"sync"
)

type ConsoleLogObserver struct {
	logger *ConsoleLogger
	mu     sync.Mutex
}

func NewConsoleLogObserver(output io.Writer) *ConsoleLogObserver {
	return &ConsoleLogObserver{
		logger: NewConsoleLogger(output),
	}
}

func (cl *ConsoleLogObserver) Update(event events.Event) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.logger.Log(0, fmt.Sprintf("[VM][EVENT][%s] - %s", event.Name, event.Description))

}

type FileLogObserver struct {
	logger *FileLogger
	mu     sync.Mutex
}

func NewFileLogObserver() *FileLogObserver {
	fl, _ := NewFileLogger("tvm.log")
	return &FileLogObserver{
		logger: fl,
	}
}

func (cl *FileLogObserver) Update(event events.Event) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.logger.Log(0, fmt.Sprintf("[VM][EVENT][%s] - %s", event.Name, event.Description))
}
