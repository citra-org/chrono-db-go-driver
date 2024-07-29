package connection

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Connection struct {
	conn net.Conn
}

func NewConnection(uri string) (*Connection, string, error) {
	parts := strings.Split(uri, "://")
	if len(parts) != 2 || parts[0] != "itlg" {
		return nil, "", fmt.Errorf("invalid URI scheme")
	}

	authAndHost := strings.Split(parts[1], "@")
	if len(authAndHost) != 2 {
		return nil, "", fmt.Errorf("invalid URI format")
	}

	auth := strings.Split(authAndHost[0], ":")
	if len(auth) != 2 {
		return nil, "", fmt.Errorf("invalid URI format")
	}
	username := auth[0]
	password := auth[1]

	hostAndDatabase := strings.Split(authAndHost[1], "/")
	if len(hostAndDatabase) != 2 {
		return nil, "", fmt.Errorf("invalid URI format")
	}
	host := hostAndDatabase[0]
	chrono := hostAndDatabase[1]

	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, "", err
	}

	connection := &Connection{conn: conn}
	err = connection.authenticate(username, password)
	if err != nil {
		conn.Close()
		return nil, "", err
	}

	return connection, chrono, nil
}

func (c *Connection) authenticate(username, password string) error {
	command := fmt.Sprintf("auth %s %s", username, password)
	_, err := c.conn.Write([]byte(command))
	if err != nil {
		return err
	}

	response, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		return err
	}

	if strings.TrimSpace(response) != "OK" {
		return fmt.Errorf("authentication failed: %s", response)
	}

	return nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}

func (c *Connection) Execute(command string) (string, error) {
	_, err := c.conn.Write([]byte(command + "\n"))
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}

	response := string(buffer[:n])
	return response, nil
}
