package redis

import (
	"fmt"

	redigo "github.com/garyburd/redigo/redis"

	"github.com/st3v/tracerr"
)

var dialer = redigo.Dial

type option func(*client)

func Password(password string) option {
	return func(c *client) {
		c.auth = password
	}
}

func Connect(host string, port int, options ...option) (Client, error) {
	conn, err := dialer("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	client := &client{
		conn: conn,
	}

	for _, option := range options {
		option(client)
	}

	err = client.authenticate()
	if err != nil {
		client.Close()
		return nil, tracerr.Wrap(err)
	}

	return client, nil
}
