package client

import (
	"fmt"
	"github.com/IntelliLog/IntelliLog-GoLang-Driver/connection"
	"strings"
)

type Client struct {
	conn *connection.Connection
}

func NewClient(uri string) (*Client, error) {
	conn, err := connection.NewConnection(uri)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Create() error {
	response, err := c.conn.Execute("c")
	if err != nil {
		return err
	}
	if response != "OK" {
		return fmt.Errorf("create failed: %s", response)
	}
	return nil
}

func (c *Client) Write(data map[string]string) error {
	command := "w"
	for k, v := range data {
		command += fmt.Sprintf(" %s %s", k, v)
	}
	response, err := c.conn.Execute(command)
	if err != nil {
		return err
	}
	if response != "OK" {
		return fmt.Errorf("write failed: %s", response)
	}
	return nil
}

func (c *Client) Read() (string, error) {
	response, err := c.conn.Execute("r")
	if err != nil {
		return "", err
	}
	fmt.Println(response)
	formattedResponse := strings.ReplaceAll(response, "&/n", "\n")
	return formattedResponse, nil
}
