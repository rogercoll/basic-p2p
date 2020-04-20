package server

import (
	"log"
	"fmt"
	"net"
	"io"
	"bytes"
	"strconv"
	"math/rand"
	"github.com/rogercoll/p2p/pkg/messages"
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

func read(conn net.Conn, least int) (*[]byte, int, error){
	log.Println("Listener: Accepted a request")
	log.Println("Listener: Read the request content...")
	buf := make([]byte, least)
	n, err := io.ReadAtLeast(conn, buf, least)
	if err != nil {
		return nil, 0, err
	}
	return &buf, n, nil
}

func handleConnection(c net.Conn) {
	//To improve performace lets use a pool of bytes for each connection
	/*
	myConnPool := &sync.Pool{
		New: func() interface{} {
			mem := make([]byte, 128)
			return &mem
		},
	}
	*/
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		content, size, err := read(c,24)
		if err == io.EOF {
                break
        } else if err != nil {
			log.Printf("Listener[ERROR]: Error while reading content %v\n", err)
		}
		if size >= 24 {
			fmt.Println(size)
			tmpbuf := bytes.NewReader(*content)
			fmt.Println(*content)
			v, err := messages.ReadVersion(tmpbuf)
			if err != nil {
				fmt.Println(err)
			}			
			readbleMessage, err := messages.Parse(v)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("+%v\n", readbleMessage)
			}
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