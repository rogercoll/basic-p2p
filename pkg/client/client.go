package client

import (
	"fmt"
	"net"
	"bufio"
	"strings"
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
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
				fmt.Println(err)
				return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
				break
		}

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