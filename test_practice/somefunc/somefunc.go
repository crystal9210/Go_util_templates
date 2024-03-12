package somefunc

type Caller interface {
	Call(val int) int
}

type Client struct {
	FuncCaller Caller
}

type ExampleCaller struct{}

func (c *Client) Run(val int) int {
	return c.FuncCaller.Call(val)
}

func (f *ExampleCaller) Call(val int) int {
	return val
}
