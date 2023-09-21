package nando

import "fmt"

type Client struct {
	Server string
}

func (c *Client) Do(req *Request) (*Response, error) {
	serving[c.Server].Submit(req)
	resp := <-serving[c.Server].responses
	fmt.Println(resp)
	return resp, nil
}
