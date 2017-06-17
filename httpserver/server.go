package httpserver

import (
	"net"
	"strconv"
	"time"
)

type server struct{
	address	string
	port	int
	conn    net.Listener
}

func StartServer(address string, port int) *server {
	return &server{address:address, port:port}
}

func (server *server) Connect() (errMsg error){
	ipAddr := server.address + ":" + strconv.Itoa(server.port)
	server.conn, errMsg = net.Listen("tcp", ipAddr)
	return
}

func (server *server) GetClient(readBuffer int, lifeTime time.Duration) (client *Client, err error) {
	clientConn, err := server.conn.Accept();
	if err != nil {
		client = nil
		return
	}

	client = NewClient(clientConn, readBuffer, lifeTime)
	return
}
