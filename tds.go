package tds

import (
	"encoding/binary"
	"io"
)

type Proxy struct {
	server *Conn
	client *Conn
}

func NewProxy(server, client io.ReadWriteCloser) *Proxy {
	p := &Proxy{
		server: newConn(server),
		client: newConn(client),
	}
	return p
}

func (p *Proxy) Handle() {
}

type Conn struct {
	transport io.ReadWriteCloser

	nextReady   chan bool
	packetReady chan bool
	errChan     chan error
}

func newConn(t io.ReadWriteCloser) *Conn {
	c := &Conn{transport: t}

	c.errChan = make(chan error)
	c.nextReady = make(chan bool)
	c.packetReady = make(chan bool)

	c.nextReady <- true

	return c
}

func (conn *Conn) readPackets() {
	for {
		<-conn.nextReady
		header := header{}
		if err := binary.Read(conn.transport, binary.BigEndian, &header); err != nil {
			conn.errChan <- err
			return
		}

		conn.packetReady <- true
	}
}

func (conn *Conn) receiverLoop() {
	for {
		select {
		case <-conn.packetReady:
			// asd
		case <-conn.errChan:
			break
		}
	}
	conn.transport.Close()
}

type header struct {
	PacketType uint8
	Status     uint8
	Size       uint16
	Spid       uint16
	Seq        uint8
	Pad        uint8
}
