package client

import (
	"fmt"
	"github.com/citra-org/chrono-db-go-driver/connection"
)

type Client struct {
	conn *connection.Connection
}

func Connect(uri string) (*Client, string, error) {
	conn, dbName, err := connection.NewConnection(uri)
	if err != nil {
		return nil, "", err
	}
	return &Client{conn: conn}, dbName, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) PingChrono() error {
	if response, err := c.conn.Execute("PING"); err != nil || response != "OK" {
		return fmt.Errorf("ping failed: %v", err)
	}
	return nil
}
func (c *Client) CreateStream(chrono string, stream string) error {
	if response, err := c.conn.Execute("CREATE STREAM " + stream); err != nil || response != "OK" {
		return fmt.Errorf("create failed: %v", err)
	}
	return nil
}
func (c *Client) DeleteStream(chrono string, stream string) error {
	if response, err := c.conn.Execute("DELETE STREAM " + stream); err != nil || response != "OK" {
		return fmt.Errorf("delete failed: %v", err)
	}
	return nil
}

func (c *Client) WriteEvent(chrono string, stream string, event string) error {
	command := "INSERT " + event + " INTO " + stream
	if response, err := c.conn.Execute(command); err != nil || response != "OK" {
		return fmt.Errorf("write failed: %v", err)
	}
	return nil
}

func (c *Client) Read(chrono string, stream string) (string, error) {
	return c.conn.Execute("SELECT * FROM " + stream)
}
