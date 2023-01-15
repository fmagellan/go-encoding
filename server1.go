package main

import (
	"encoding/binary"
	"fmt"
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

	b := make([]byte, 1024)
mainloop:
	for i := 0; i < 2; i++ {
		n, err := conn.Read(b)
		if err != nil {
			panic(err)
		}

		op := string(b[:n])
		fmt.Printf("op: %v\n", op)
		switch op {
		case "FixedStructWithReflection":
			sFixedStructWithReflection(conn)
		case "StructWithVarInt":
			sStructWithVarInt(conn)
		default:
			fmt.Printf("invalid op: %v\n", op)
			break mainloop
		}
	}
}

func sFixedStructWithReflection(conn net.Conn) {
	s := utils.Numbers{
		N1: 10,
		N2: 11,
		N3: 12,
	}

	err := binary.Write(conn, binary.BigEndian, s)
	if err != nil {
		panic(err)
	}
}

func sStructWithVarInt(conn net.Conn) {
	s := utils.Numbers{
		N1: 10,
		N2: 11,
		N3: 12,
	}

	b := make([]byte, 1024)
	n := binary.PutVarint(b, int64(s.N1))
	n += binary.PutVarint(b[n:], int64(s.N2))
	n += binary.PutVarint(b[n:], int64(s.N3))
	_, err := conn.Write(b[:n])
	if err != nil {
		panic(err)
	}
}
