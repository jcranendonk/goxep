package goxep

import (
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func Open(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	c := &Client{conn: conn}
	if err := c.initStream(); err != nil {
		go conn.Close()
		return nil, err
	}

	return c, nil
}

func (c *Client) initStream() error {
	fmt.Printf(streamStart, "a", "b")
	return nil
}

func (c *Client) negotiate() error {
	return nil
}
