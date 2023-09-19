package nando

type Request struct {
	funcName string
	data     interface{}
}

func NewRequest(funcName string, data interface{}) *Request {
	return &Request{funcName: funcName, data: data}
}
