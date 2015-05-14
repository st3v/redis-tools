package redis

import redigo "github.com/garyburd/redigo/redis"

type Client interface {
	Snapshot(doneChan chan struct{}, errChan chan error)
	Close() error
}

type client struct {
	conn redigo.Conn
	auth string
}

func (c *client) authenticate() error {
	var err error

	if c.auth != "" {
		_, err = c.conn.Do("AUTH", c.auth)
	}

	return err
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Snapshot(doneChan chan struct{}, errChan chan error) {
	// todo
}
