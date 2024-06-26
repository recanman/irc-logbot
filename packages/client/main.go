package client

import (
	"fmt"

	irc "github.com/fluffle/goirc/client"
)

type ClientOptions struct {
	Channels []string
}

type Client struct {
	Conn          *irc.Conn
	ClientOptions ClientOptions
}

func Create(server string, port int, nickname string, options ClientOptions) *Client {
	config := irc.NewConfig(nickname)
	config.SSL = false
	config.Server = fmt.Sprintf("%s:%d", server, port)
	config.NewNick = func(n string) string { return n + "^" }

	client := irc.Client(config)
	client.EnableStateTracking()

	client.HandleFunc(irc.CONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			for _, channel := range options.Channels {
				conn.Join(channel)
			}
		})

	return &Client{
		Conn:          client,
		ClientOptions: options,
	}
}

func (c *Client) Connect() error {
	if err := c.Conn.Connect(); err != nil {
		return err
	}

	return nil
}
