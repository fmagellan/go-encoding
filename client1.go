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

	s := new(utils.Numbers)
	err = binary.Read(c, binary.BigEndian, s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %v\n", s)
}
