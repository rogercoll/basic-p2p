# Basic p2p

Basic golang code to establish a TCP connection and exchange information

Example of the struct that is shared:

```go
type Version struct {
	version	        uint32
	timestamp	uint64
	addr_recv	[4]uint8 
	addr_recv_port  uint16	
	addr_trans	[4]uint8
	addr_trans_port uint16
}
```

All is encoden in Little Endian order. 

## Run

```sh
//Launch the server
go run main.go


//Create a client
go test pkg/client/client_test.go
```
