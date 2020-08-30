package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new Client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "an",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) nick(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("Nickname %s Set", c.nick))
}

func (s *server) join(c *client, args []string) {
	roomname := args[1]

	r, ok := s.rooms[roomname]

	if !ok {
		r = &room{
			name:    roomname,
			members: make(map[net.Addr]*client),
		}

		s.rooms[roomname] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c, args)

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.nick))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
}

func (s *server) listRooms(c *client, args []string) {

	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("availaible rooms are: %s", strings.Join(rooms, ", ")))

}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must join a room first"))
		return
	}

	msg := strings.Join(args[1:len(args)], " ")
	c.room.broadcast(c, c.nick+": "+msg)
}

func (s *server) quit(c *client, args []string) {
	log.Printf("client has dosconnected %s:", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c, args)

	c.msg("Quit Successfull")
}

func (s *server) quitCurrentRoom(c *client, args []string) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}
