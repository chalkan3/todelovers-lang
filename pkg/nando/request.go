package nando

type Request struct {
	funcName string
	Data     interface{}
}

func NewRequest(funcName string, data interface{}) *Request {
	return &Request{funcName: funcName, Data: data}
}
