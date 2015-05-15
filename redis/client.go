package redis

import (
	"strings"

	redigo "github.com/garyburd/redigo/redis"
)

type Client interface {
	Snapshot(doneChan chan struct{}, errChan chan error)
	Close() error
}

type client struct {
	conn    redigo.Conn
	auth    string
	aliases map[string]string
}

func newClient(conn redigo.Conn) *client {
	return &client{
		conn:    conn,
		aliases: map[string]string{},
	}
}

func (c *client) authenticate() error {
	var err error

	if c.auth != "" {
		_, err = c.conn.Do("AUTH", c.auth)
	}

	return err
}

func (c *client) lookupAlias(cmd string) string {
	alias, found := c.aliases[strings.ToUpper(cmd)]
	if !found {
		return cmd
	}
	return alias
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Snapshot(doneChan chan struct{}, errChan chan error) {
	// todo
}
