package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	initialMessage    = "[Debugger IS[%v], Mode[%s]]"
	registeredMessage = "[Log Register on line %s and has %v references]"
)

type DebugMode int

const (
	Stack = iota
	JSON
	YAML
	Raw
)

func (dm DebugMode) String() string {
	return [...]string{"Stack", "YAML", "Raw"}[dm]
}

type exec func()
type DebuggerConfig struct {
	Enable bool
	Mode   DebugMode
	Exec   exec
}

type Debugger struct {
	config    *DebuggerConfig
	reference int
}

func NewDebugger(config *DebuggerConfig) *Debugger {
	fmt.Printf(initialMessage, config.Enable, config.Mode.String())
	return &Debugger{
		config: config,
	}
}

func (d *Debugger) increaseReference() { d.reference++ }
func (d *Debugger) Register(line string) {
	d.increaseReference()
	fmt.Printf(registeredMessage, line, d.reference)

}

func (d *Debugger) debug(node Node, indent string) {
	fmt.Printf("%sType: %s, Token: %v\n", indent, node.Type().String(), node.Token())

	if isNewContext(node.Type()) {
		for _, child := range node.(*FunctionCallNode).Arguments {
			d.debug(child, indent+"  ")
		}
	}
}

func (d *Debugger) printJSON(input interface{}) {
	bb, _ := json.Marshal(input)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, bb, "", "\t")
	fmt.Println(string(prettyJSON.Bytes()))
}

func (d *Debugger) Log(node *Node) {
	if d.config.Enable {
		switch d.config.Mode {
		case Stack:
			d.debug(*node, "")
		case YAML:
			d.printJSON(node)
		case JSON:
			d.printJSON(node)
		default:
			fmt.Println(node)

		}
	}

}
