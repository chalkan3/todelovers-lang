package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type LoggerFn func(input ...interface{})
type DebugMode int

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

	fmt.Printf("%sType: %s, Token: %v\n", indent, node.Type().String(), node.Token())

	if isNewContext(node.Type()) {
		for _, child := range node.(*FunctionCallNode).Arguments {
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

func (dm DebugMode) String() string {
	return [...]string{"Stack", "JSON", "YAML", "Raw"}[dm]
}

func (dm DebugMode) Mode() LoggerFn {
	return [...]LoggerFn{printStack, printJSON, printJSON, printJSON}[dm]
}

type LoggerConfig struct {
	Enable     bool
	Mode       DebugMode
	BufferSize int
}

type Logger struct {
	logChannel chan []interface{}
	config     *LoggerConfig
}

func NewLogger(config *LoggerConfig) *Logger {
	logger := &Logger{
		logChannel: make(chan []interface{}, config.BufferSize),
		config:     config,
	}
	go logger.logWorker()
	return logger
}

func (l *Logger) logWorker() {
	if l.config.Enable {
		fn := l.config.Mode.Mode()
		for msg := range l.logChannel {
			fn(msg)
		}
	}
}

func (l *Logger) Log(message ...interface{}) {
	select {
	case l.logChannel <- message:
	default:
		fmt.Println("Logger channel is full, dropping log message:", message)
	}
}
