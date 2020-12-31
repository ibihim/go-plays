// Example from https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
package main

import (
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	workloadBufferSize = 16
)

// Mux is an actor that sends messages to different connections.
type Mux struct {
	ops chan func(map[net.Addr]net.Conn)
}

// New returns a Mux with a default workload buffer of size 16.
// 16 due to https://youtu.be/yCbon_9yGVs.
func New() *Mux {
	ops := make(chan func(map[net.Addr]net.Conn), workloadBufferSize)

	return &Mux{
		ops: ops,
	}
}

// Add tells the actor to add a connection.
func (mx *Mux) Add(cnn net.Conn) {
	mx.ops <- func(m map[net.Addr]net.Conn) {
		m[cnn.RemoteAddr()] = cnn
	}
}

// Remove tells the actor to remove a connection.
func (mx *Mux) Remove(addr net.Addr) {
	mx.ops <- func(m map[net.Addr]net.Conn) {
		delete(m, addr)
	}
}

// SendMsg sends a message to the connections.
func (mx *Mux) SendMsg(msg string) error {
	result := make(chan error, 1)

	mx.ops <- func(m map[net.Addr]net.Conn) {
		for _, cnn := range m {
			if _, err := io.WriteString(cnn, msg); err != nil {
				result <- err
				return
			}
		}

		result <- nil
	}

	return <-result
}

// TryToSendMsg tries to send a message immediately. If the message would end
// up in a queue, return an error.
func (mx *Mux) TryToSendMsg(msg string) error {
	result := make(chan error, 1)

	// select on the input to a chan to catch full workload buffers.
	select {
	case mx.ops <- func(m map[net.Addr]net.Conn) {
		for _, cnn := range m {
			if _, err := io.WriteString(cnn, msg); err != nil {
				result <- err
				return
			}
		}

		result <- nil
	}:
	default:
		result <- errors.New("Too man messages")
	}

	return <-result
}

// PrivateMsg sends a private message and has the capability to return an
// error.
func (mx *Mux) PrivateMsg(addr net.Addr, msg string) error {
	result := make(chan net.Conn, 1)
	mx.ops <- func(m map[net.Addr]net.Conn) {
		result <- m[addr]
	}

	cnn := <-result
	if cnn == nil {
		return fmt.Errorf("client %v not registered", addr)
	}

	_, err := io.WriteString(cnn, msg)
	if err != nil {
		return err
	}

	return nil
}

// loop starts the actor.
func (mx *Mux) loop() {
	cnns := make(map[net.Addr]net.Conn)

	for op := range mx.ops {
		op(cnns)
	}
}
