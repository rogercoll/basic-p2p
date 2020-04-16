package client

import (
	"log"
	"net"
	"github.com/rogercoll/p2p/pkg/messages"
	"bytes"
	"reflect"
	"encoding/binary"
)

const (
	PORT = "7777"
	PROTOCOL = "tcp"
	ADDRESS = "127.0.0.1"
)

func WriteMessage(c net.Conn, message string) {
	buf := new(bytes.Buffer)
	data := messages.Version{
		uint32(1),
		uint64(12345),
		[4]byte{192,168,1,1},
		uint16(123),
		[4]byte{192,168,1,48},
		uint16(123),
	}
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		err := binary.Write(buf, binary.LittleEndian, v.Field(i).Interface())
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
	}
	log.Printf("%x", buf.Bytes())
	log.Println(len(buf.Bytes()))
	log.Println(message)
	c.Write(buf.Bytes())
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