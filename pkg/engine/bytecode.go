package engine

import (
	"fmt"
	interpreter "mary_guica/pkg/tvm/pkg/runtime"
)

func PrintByteCode(code []byte) {

	format := map[byte]func(b []byte, pos int, sFormat *string) int{
		interpreter.ADD: func(b []byte, pos int, sFormat *string) int {
			command := b[pos : pos+2]
			*sFormat += fmt.Sprintf("%s %d %d \n", "ADD", command[0], command[1])

			return pos + 3
		},
		interpreter.PRINT: func(b []byte, pos int, sFormat *string) int {
			*sFormat += fmt.Sprintf("%s %d \n", "PRINT", b[pos+1])
			return pos + 2
		},
		interpreter.MOV: func(b []byte, pos int, sFormat *string) int {
			command := b[pos : pos+2]
			*sFormat += fmt.Sprintf("%s %d %d \n", "MOV", command[0], command[1])
			return pos + 3
		},
		interpreter.STORE: func(b []byte, pos int, sFormat *string) int {
			command := b[pos : pos+2]
			*sFormat += fmt.Sprintf("%s %d %d \n", "STORE", command[0], command[1])
			return pos + 3
		},
		interpreter.HALT: func(b []byte, pos int, sFormat *string) int {
			*sFormat += "HALT\n"
			return pos + 1
		},
		interpreter.LOAD_STRING: func(b []byte, pos int, sFormat *string) int {
			len := b[pos+2]

			fa := []interface{}{"LOAD_STRING", b[pos+1], len, string(b[pos+3 : pos+3+int(len)])}

			*sFormat += fmt.Sprintf("%s %d %d %s \n", fa...)
			return pos + int(len) + 3
		},
		interpreter.CREATE_VAR: func(b []byte, pos int, sFormat *string) int {
			len := b[pos : pos+1][0]

			fa := []interface{}{"CREATE_VAR", len, string(b[pos+1 : pos+1+int(len)])}

			*sFormat += fmt.Sprintf("%s %d %s \n", fa...)
			return pos + int(len) + 2
		},
		interpreter.LOAD_FROM_VAR: func(b []byte, pos int, sFormat *string) int {
			len := b[pos : pos+1][0]
			fa := []interface{}{"LOAD_FROM_VAR", b[pos : pos+1][0], string(b[pos+1 : pos+1+int(len)])}

			*sFormat += fmt.Sprintf("%s %d %s \n", fa...)
			return pos + int(len) + 2
		},
		interpreter.LOAD: func(b []byte, pos int, sFormat *string) int {
			*sFormat += fmt.Sprintf("%s %d \n", "LOAD", b[pos+1])
			return pos + 2
		},
		interpreter.S_THREAD: func(b []byte, pos int, sFormat *string) int {
			*sFormat += fmt.Sprintf("%s %d \n", "S_THREAD", b[pos+1])
			return pos + 2

		},
		interpreter.ST_THREAD: func(b []byte, pos int, sFormat *string) int {
			*sFormat += "ST_THREAD\n"
			return pos + 1

		},
		interpreter.W_THREAD: func(b []byte, pos int, sFormat *string) int {
			*sFormat += "W_THREAD\n"
			return pos + 1

		},
		interpreter.SY_THREAD: func(b []byte, pos int, sFormat *string) int {
			*sFormat += "SY_THREAD\n"
			return pos + 1

		},
		interpreter.CALL: func(b []byte, pos int, sFormat *string) int {
			*sFormat += fmt.Sprintf("%s %d \n", "CALL", b[pos+1])
			return pos + 2
		},
		interpreter.RET: func(b []byte, pos int, sFormat *string) int {
			*sFormat += "RET\n"
			return pos + 1
		},
	}
	formattedString := ""

	curr := 0
	for {
		if curr == len(code) {
			break
		}
		a := code[curr]
		_ = a
		curr = format[byte(code[curr])](code, curr, &formattedString)
	}

	// Imprimir a string formatada
	fmt.Println(formattedString)
}
