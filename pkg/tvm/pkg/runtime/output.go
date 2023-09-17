package runtime

import (
	"mary_guica/pkg/cast"
	"mary_guica/pkg/tvm/pkg/register"
)

type Output struct {
	v interface{}
}

func (o Output) ToByte() byte        { return o.v.(byte) }
func (o Output) ToByteArray() []byte { return o.v.([]byte) }

func (o Output) ToInt() int                    { return cast.ToAlwaysInt(o.v) }
func (o Output) ToString() string              { return o.v.(string) }
func (o Output) ToFloat() string               { return o.v.(string) }
func (o Output) ToRegister() register.Register { return o.v.(register.Register) }
