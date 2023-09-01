package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type loggerFn func(input ...interface{})
type debugMode int

const (
	Stack = iota
	JSON
	YAML
	Raw
)

func printStack(input ...interface{}) {
	if len(input) < 2 {
		panic("Logger need 2 params")
	}

	node := input[0].(Node)
	indent := input[1].(string)

	fmt.Printf("%sType: %s, token: %v\n", indent, node.Type().String(), node.Token())

	if isNewContext(node.Type()) {
		for _, child := range node.(*functionCallNode).Arguments {
			printStack(child, indent+"  ")
		}
	}
}

func printJSON(input ...interface{}) {
	bb, _ := json.Marshal(input[0])
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, bb, "", "\t")
	fmt.Println(string(prettyJSON.Bytes()))
}

func (dm debugMode) String() string {
	return [...]string{"Stack", "JSON", "YAML", "Raw"}[dm]
}

func (dm debugMode) Mode() loggerFn {
	return [...]loggerFn{printStack, printJSON, printJSON, printJSON}[dm]
}

type LoggerConfig struct {
	Enable     bool
	Mode       debugMode
	BufferSize int
}

type logger struct {
	logChannel chan []interface{}
	config     *LoggerConfig
}

func NewLogger(config *LoggerConfig) *logger {
	logger := &logger{
		logChannel: make(chan []interface{}, config.BufferSize),
		config:     config,
	}
	go logger.logWorker()
	return logger
}

func (l *logger) logWorker() {
	if l.config.Enable {
		fn := l.config.Mode.Mode()
		for msg := range l.logChannel {
			fn(msg)
		}
	}
}

func (l *logger) Log(message ...interface{}) {
	select {
	case l.logChannel <- message:
	default:
		fmt.Println("Logger channel is full, dropping log message:", message)
	}
}
