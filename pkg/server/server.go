package server

import (
	"fmt"
	"net"
	"bufio"
	"bytes"
	"io"
	"strconv"
	"math/rand"
)

const (
	PORT = "7777"
	PROTOCOL = "tcp"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func newServer() (net.Listener, error) {
	l, err := net.Listen(PROTOCOL, ":" + PORT)
	if err != nil {
			l.Close()
			return nil, err
	}
	return l, nil
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	reader := bufio.NewReader(c)
	var buffer bytes.Buffer

	for {
		netData,  isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return 
		}
		buffer.Write(netData)
		temp := buffer.String()
		if temp == "STOP" {
				break
		}
		if !isPrefix {
			break
		}
		fmt.Println("yeeeee")
		fmt.Println(temp)

		result := strconv.Itoa(random(1,1234)) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()
}


func Run() error {
	ln, err := newServer()
	defer ln.Close()
	if err != nil {
		return err
	}

	for {
		c, err := ln.Accept()
		if err != nil {
				return err
		}
		go handleConnection(c)
	}
}