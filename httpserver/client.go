package httpserver

import (
	"net"
	"fmt"
	"io"
	"time"
)

type Client struct {
	Conn net.Conn
	readBuffer int
	lifeTime time.Duration
	WaitChan chan *Request  /*map[string]string*/
}

type HttpHandle interface {
	HandleMethod(request *Request) ([]byte, error)
}

func NewClient(conn net.Conn, readBuffer int, lifeTime time.Duration) (client *Client) {
	//connChan := make(chan map[string]string)
	connChan := make(chan *Request)
	client = &Client{conn, readBuffer, lifeTime, connChan}
	client.Conn.SetReadDeadline(time.Now().Add(time.Second * lifeTime))

	return
}

type ParseRequest interface {
	ParseRequest(requestBytes []byte) map[string]interface{}
}

func (client *Client) GetRequest(request ParseRequest) {
	var requestBytes []byte
	requestByte := make([]byte, client.readBuffer)

	for {
		len, err := client.Conn.Read(requestByte)
		if err != nil {
			if err == io.EOF  {
				break
			}
			fmt.Println(err)
		}

		requestBytes = append(requestBytes, requestByte[:len]...)

		if (len < client.readBuffer) {
			client.Conn.SetReadDeadline(time.Time{})
			break
		}
	}

	requestData := request.ParseRequest(requestBytes)
	requestStruct := NewRequest(requestData)

	client.WaitChan <- requestStruct
}

func (client *Client) SetResponse(handle HttpHandle) {
	defer func() {
		client.Conn.Close()
		close(client.WaitChan)
	}()

	requestStr := <-client.WaitChan
	content, err :=handle.HandleMethod(requestStr)
	_, err = client.Conn.Write(content)
	if err != nil {
		fmt.Println(err, requestStr)
	}
}
