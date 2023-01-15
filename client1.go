package main

import (
	"encoding/binary"
	"fmt"
	"github.com/fmagellan/go-encoding/utils"
	"net"
)

func main() {
	c, err := net.Dial("tcp", utils.ServerPort)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	cFixedStructWithReflection(c)
	cStructWithVarInt(c)
}

func cFixedStructWithReflection(c net.Conn) {
	_, _ = c.Write([]byte("FixedStructWithReflection"))

	s := new(utils.Numbers)
	err := binary.Read(c, binary.BigEndian, s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %v\n", s)
}

func cStructWithVarInt(c net.Conn) {
	_, _ = c.Write([]byte("StructWithVarInt"))

	b := make([]byte, 1024)
	n, err := c.Read(b)
	if err != nil {
		panic(err)
	}

	n1, n1Bytes := binary.Varint(b)
	n2, n2Bytes := binary.Varint(b[n1Bytes:])
	n3, n3Bytes := binary.Varint(b[n1Bytes+n2Bytes:])

	s := utils.Numbers{
		N1: int32(n1),
		N2: int8(n2),
		N3: uint16(n3),
	}
	fmt.Printf("number of bytes read: %v, var int bytes decoded: %v\n", n, n1Bytes+n2Bytes+n3Bytes)
	fmt.Printf("data: %v\n", s)
}
