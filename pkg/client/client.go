package client

import (
	"log"
	"net"
	"github.com/rogercoll/p2p/pkg/messages"
	"bytes"
)

const (
	PORT = "7777"
	PROTOCOL = "tcp"
	ADDRESS = "127.0.0.1"
)

func WriteMessage(c net.Conn, message string) {
	data, err := messages.NewVersion(12, []byte{192,168,1,1},123, []byte{192,168,1,48}, 775)
	if err != nil {
		log.Println(err)
	}
	var b bytes.Buffer
	n, err := data.WriteTo(&b)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%x", b.Bytes())
	log.Println(n)
	log.Println(*data)
	c.Write(b.Bytes())
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