package server

import (
	"log"
	"net"
)

const (
	PORT = "7777"
	PROTOCOL = "tcp"
	ADDRESS = "127.0.0.1"
)

func WriteMessage(c net.Conn, message string) {
	c.Write([]byte(message))
}

func Run() error {
	conn, err := net.Dial("tcp", ADDRESS + ":" + PORT)
    if err != nil {
		return err
	}
	defer conn.Close()
	WriteMessage(conn, "hello bitch\n")
	reply := make([]byte, 1024)
	conn.Read(reply)
	log.Println(string(reply))
	return nil
}