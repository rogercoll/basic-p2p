package messages

import (
	"time"
)

type Message interface {
	//Parse...
}
//The version message provides information about the transmitting node to the receiving node at the beginning of a connection. 
//Until both peers have exchanged version messages, no other messages will be accepted.

type Version struct {
	Version 	uint32
	Timestamp	uint64
	Addr_recv 	[4]byte //The IPv6 address of the receiving node as perceived by the transmitting node in big endian byte order. IPv4 addresses can be provided as IPv4-mapped IPv6 addresses.
	Addr_recv_port uint16	//Big Endian byte order
	Addr_trans	[4]byte
	Addr_trans_port uint16
}

//Bytes= 4 + 8 + 16 + 2 + 16 + 2 = 48 bytes

type VersionReadble struct {
	Version 	string
	Timestamp	time.Time
}