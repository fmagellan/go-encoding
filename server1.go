package main

import (
	"encoding/binary"
	"github.com/fmagellan/go-encoding/utils"
	"net"
)

func main() {
	l, err := net.Listen("tcp", utils.ServerPort)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	s := utils.Numbers{
		N1: 10,
		N2: 11,
		N3: 12,
	}

	err = binary.Write(conn, binary.BigEndian, s)
	if err != nil {
		panic(err)
	}
}
