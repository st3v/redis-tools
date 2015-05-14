package fake

type Conn struct {
	receivedDoCalls map[string][]interface{}
	ExpectedDoErr   error
	ExpectedDoReply interface{}

	receivedCloseCalls int
	ExpectedCloseErr   error
}

func (c *Conn) ReceivedDoCalls() map[string][]interface{} {
	return c.receivedDoCalls
}

func (c *Conn) ReceivedCloseCalls() int {
	return c.receivedCloseCalls
}

func (c *Conn) Close() error {
	c.receivedCloseCalls++
	return c.ExpectedCloseErr
}

func (c *Conn) Err() error {
	return nil
}

func (c *Conn) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	if c.receivedDoCalls == nil {
		c.receivedDoCalls = map[string][]interface{}{}
	}
	c.receivedDoCalls[cmd] = args
	return c.ExpectedDoReply, c.ExpectedDoErr
}

func (c *Conn) Send(cmd string, args ...interface{}) error {
	return nil
}

func (c *Conn) Flush() error {
	return nil
}

func (c *Conn) Receive() (reply interface{}, err error) {
	return nil, nil
}
