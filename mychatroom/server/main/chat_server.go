// @author     : wyettLei
// @date       : Created in 2021/4/20 15:47
// @description: TODO

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mychatroom/protocol"
	"net"
	"sync"
)

type ChatServer interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Start()
	Close()
}

type client struct {
	conn   net.Conn
	name   string
	writer *protocol.CommandWriter
}

type TcpChatServer struct {
	listener net.Listener
	clients  []*client
	m        *sync.Mutex
}

func (tcs *TcpChatServer) SetListener(listener net.Listener) {
	tcs.listener = listener
}

// create listener
func (tcs *TcpChatServer) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errors.New(fmt.Sprintf("listen to %v failed\n", address))
	}
	tcs.SetListener(listener)
	log.Printf("Listen to %v", address)
	return nil
}

func (tcs *TcpChatServer) Broadcast(command interface{}) error {
	for _, client := range tcs.clients {
		client.writer.Write(command)
	}
	return nil
}

func (tcs *TcpChatServer) Start() {
	for {
		conn, err := tcs.listener.Accept()
		if err != nil {
			log.Print(err)
		} else {
			client := tcs.accept(conn)
			go tcs.serve(client)
		}
	}
}

func (tcs *TcpChatServer) Close() {
	tcs.listener.Close()
}

func (tcs *TcpChatServer) serve(client *client) {
	cmdReader := protocol.NewCommandReader(client.conn)
	defer tcs.remove(client)
	for {
		cmd, err := cmdReader.Read()
		if err != nil && err != io.EOF {
			log.Printf("Read error: %v", err)
		}
		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.SendCommand:
				go tcs.Broadcast(protocol.MessageCommand{
					Message: v.Message,
					Name:    client.name,
				})
			case protocol.NameCommand:
				client.name = v.Name
			}
		}
		if err == io.EOF {
			break
		}
	}
}

func (tcs *TcpChatServer) accept(conn net.Conn) *client {
	tcs.m.Lock()
	defer tcs.m.Unlock()
	client := &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	}
	tcs.clients = append(tcs.clients, client)
	return client
}

func (tcs *TcpChatServer) remove(client *client) {
	tcs.m.Lock()
	defer tcs.m.Unlock()
	for i, check := range tcs.clients {
		if check == client {
			tcs.clients = append(tcs.clients[:i], tcs.clients[i+1:]...)
		}
	}
	log.Printf("close connection from %v", client.conn.RemoteAddr().String())
	client.conn.Close()
}
