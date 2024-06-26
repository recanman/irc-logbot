package ircbot

import (
	"fmt"
	"strings"

	irc "github.com/fluffle/goirc/client"
	"github.com/recanman/irc-logbot/packages/client"
	"github.com/recanman/irc-logbot/packages/logger"
)

func strip(s string) string {
	s = strings.Map(func(r rune) rune {
		if r == '\x01' {
			return -1
		}
		return r
	}, s)
	return s
}

func handleEvents(client *client.Client, eventLogger *logger.EventLogger) {
	client.Conn.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
		for _, channel := range client.ClientOptions.Channels {
			conn.Join(channel)
		}
	})

	client.Conn.HandleFunc(irc.JOIN, func(conn *irc.Conn, line *irc.Line) {
		eventLogger.LogEvent(line.Args[0], fmt.Sprintf("%s joined.", line.Nick))
	})

	client.Conn.HandleFunc(irc.PART, func(conn *irc.Conn, line *irc.Line) {
		eventLogger.LogEvent(line.Args[0], fmt.Sprintf("%s left (%s).", line.Nick, line.Args[1]))
	})

	client.Conn.HandleFunc(irc.KICK, func(conn *irc.Conn, line *irc.Line) {
		eventLogger.LogEvent(line.Args[0], fmt.Sprintf("%s was kicked by %s (%s).", line.Args[1], line.Nick, line.Args[2]))
	})

	client.Conn.HandleFunc(irc.TOPIC, func(conn *irc.Conn, line *irc.Line) {
		eventLogger.LogEvent(line.Args[0], fmt.Sprintf("Topic is \"%s\" (set by %s).", line.Args[1], line.Nick))
	})

	client.Conn.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		eventLogger.LogEvent(line.Args[0], fmt.Sprintf("<%s> %s", line.Nick, strip(line.Args[1])))
	})
}

func FromClient(client *client.Client, fileName string) {
	eventLogger, err := logger.CreateEventLogger(fileName)
	if err != nil {
		panic(err)
	}

	handleEvents(client, eventLogger)
}
