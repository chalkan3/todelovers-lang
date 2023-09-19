package nando

type Client struct{}

func (c Client) Do(req *Request) (*Response, error) {
	serving.Submit(req)
	resp := <-serving.responses
	return resp, nil
}
