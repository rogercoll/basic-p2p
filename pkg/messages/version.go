package messages

import (
	"io"
	"fmt"
	"time"
	"bytes"
	//"reflect"
	"encoding/binary"
)

type Message interface {
	//Parse...
}
//The version message provides information about the transmitting node to the receiving node at the beginning of a connection. 
//Until both peers have exchanged version messages, no other messages will be accepted.

type Version struct {
	version	        uint32
	timestamp	uint64
	addr_recv	[]byte //The IPv6 address of the receiving node as perceived by the transmitting node in big endian byte order. IPv4 addresses can be provided as IPv4-mapped IPv6 addresses.
	addr_recv_port  uint16	//Big Endian byte order
	addr_trans	[]byte
	addr_trans_port uint16
}
//Bytes= 4 + 8 + 16 + 2 + 16 + 2 = 48 bytes


func NewVersion(v int, addr_recv []byte, recv_port int, addr_trans []byte, trans_port int) (*Version, error) {
	if len(addr_recv) != 4 {
		return nil, fmt.Errorf("Invalid reciver address: %v\n",addr_recv )
	} else if len(addr_trans) != 4 {
		return nil, fmt.Errorf("Invalid transmitting address: %v\n", addr_trans )
	}
	newObject := Version{
		uint32(v),
		uint64(time.Now().Unix()),
		addr_recv,
		uint16(recv_port),
		addr_trans,
		uint16(trans_port),
	}
	return &newObject, nil
}


func (data *Version) WriteTo(w io.Writer) (int64, error) {
	buf := new(bytes.Buffer)
	for _, d := range []interface{}{data.version, data.timestamp, data.addr_recv, data.addr_recv_port, data.addr_trans, data.addr_trans_port} {
		err := binary.Write(buf, binary.LittleEndian, d)
		if err != nil {
			return 0, fmt.Errorf("binary.Write failed: %v\n", err)
		}
	}
	n, err := w.Write(buf.Bytes())
	if err != nil {
		return int64(n), fmt.Errorf("There was an error while writing into io.Writer: %v\n", err)
	}
	return int64(n), nil
}

func ReadVersion(r io.Reader) (*Version, error) {
	var v Version
	err := binary.Read(r, binary.LittleEndian, &v)
	if err != nil {
		fmt.Println("hereeee")
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	return &v, nil
}


type VersionReadble struct {
	Version	        string
	Timestamp	time.Time
}
