package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// Arith
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

type mathConn struct {
	net.Conn
	readBuf  bytes.Buffer
	writeBuf bytes.Buffer
	A, B     int
	Quo, Rem int
	closec   chan bool
}

func (c *mathConn) Read(b []byte) (int, error) {
	return c.readBuf.Read(b)
}

func (c *mathConn) Write(b []byte) (int, error) {
	return c.writeBuf.Write(b)
}

func (c *mathConn) Close() error {
	select {
	case c.closec <- true:
	default:
	}
	return nil
}

func handleConnection(c net.Conn) {
	fmt.Println("Got a connection ", c)
	c.Write(bytes.NewBufferString("HELLO!!").Bytes())
}

func main() {
	rpc.Register(new(Arith))

	// listen on the port
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listening...", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		go handleConnection(conn)
	}
}
