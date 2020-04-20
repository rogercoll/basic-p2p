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
	addr_recv	[4]uint8 //The IPv6 address of the receiving node as perceived by the transmitting node in big endian byte order. IPv4 addresses can be provided as IPv4-mapped IPv6 addresses.
	addr_recv_port  uint16	//Big Endian byte order
	addr_trans	[4]uint8
	addr_trans_port uint16
}

func NewVersion(v int, addr_recv []byte, recv_port int, addr_trans []byte, trans_port int) (*Version, error) {
	if len(addr_recv) != 4 {
		return nil, fmt.Errorf("Invalid reciver address: %v\n",addr_recv )
	} else if len(addr_trans) != 4 {
		return nil, fmt.Errorf("Invalid transmitting address: %v\n", addr_trans )
	}
	var addr_recv_uint8 [4]uint8
	for i, adr := range addr_recv {
		addr_recv_uint8[i] = uint8(adr)
	}
	var addr_trans_uint8 [4]uint8
	for i, adr := range addr_trans {
		addr_trans_uint8[i] = uint8(adr)
	}
	newObject := Version{
		uint32(v),
		uint64(time.Now().Unix()),
		addr_recv_uint8,
		uint16(recv_port),
		addr_trans_uint8,
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
	fmt.Printf("+%v\n",r)
	var v Version
	err := binary.Read(r, binary.LittleEndian, &v.version)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	err = binary.Read(r, binary.LittleEndian, &v.timestamp)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	err = binary.Read(r, binary.LittleEndian, &v.addr_recv)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	err = binary.Read(r, binary.LittleEndian, &v.addr_recv_port)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	err = binary.Read(r, binary.LittleEndian, &v.addr_trans)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	err = binary.Read(r, binary.LittleEndian, &v.addr_trans_port)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	return &v, nil
}

//Bytes= 4 + 8 + 16 + 2 + 16 + 2 = 48 bytes

type VersionReadble struct {
	Version	        string
	Timestamp	time.Time
}
