package cast

import (
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/register"
	"mary_guica/pkg/tvm/pkg/threads"
)

func ToInt(cast interface{}) int { return cast.(int) }
func ToUInt8(cast interface{}) (uint8, bool) {
	c, ok := cast.(uint8)
	return c, ok
}

func ToByte(cast interface{}) byte                             { return cast.(byte) }
func ToByteArray(cast interface{}) []byte                      { return cast.([]byte) }
func ToProgramManager(cast interface{}) program.ProgramManager { return cast.(program.ProgramManager) }
func ToThreadManager(cast interface{}) threads.ThreadManager   { return cast.(threads.ThreadManager) }
func ToMemoryManager(cast interface{}) memory.MemoryManager    { return cast.(memory.MemoryManager) }
func ToRegisterManager(cast interface{}) register.RegisterManager {
	return cast.(register.RegisterManager)
}

func ToRegister(cast interface{}) register.Register { return cast.(register.Register) }

func ToString(cast interface{}) (string, bool) {
	s, err := cast.(string)
	return s, err
}

func ToAlwaysInt(cast interface{}) int {
	castUInt8, ok := ToUInt8(cast)
	if ok {
		return int(castUInt8)
	}

	return ToInt(cast)
}
