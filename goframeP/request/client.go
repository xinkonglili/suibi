package request

import "sync"

var registry = make(map[string]CommonClient)

func register(c CommonClient) {
	registry[c.GetClientCode()] = c
}

func init() {
	register(&Client{flat: "client"})
	register(&ClientA{flat: "clientA"})
}

type CommonClient interface {
	GetClientCode() string
	Post(ctx, url string, body []byte, options ...Options) ([]byte, error)
}
type Client struct {
	mu   sync.RWMutex
	flat string
}

type ClientA struct {
	mu   sync.RWMutex
	flat string
}

type Options struct {
	Method  string
	Timeout int // 请求超时时间，单位：毫秒
	Headers map[string]string
}

func (c *Client) GetClientCode() string {
	return c.flat
}

func (c *Client) Post(ctx, url string, body []byte, options ...Options) ([]byte, error) {
	return nil, nil
}

func (c *ClientA) GetClientCode() string {
	return c.flat
}

func (c *ClientA) Post(ctx, url string, body []byte, options ...Options) ([]byte, error) {

	return nil, nil
}
