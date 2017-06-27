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

func (client *Client) SetReponse(handle HttpHandle) {
	defer func() {
		client.Conn.Close()
		close(client.WaitChan)
	}()

	requestStr := <-client.WaitChan
	fmt.Println(requestStr)
	content, err :=handle.HandleMethod(requestStr)
	responseStr := "HTTP/1.1 200 OK\r\n"
	responseStr += "Expires:Tue, 13 Jun 2017 11:57:00 GMT\r\n"
	responseStr += "Content-Type:text/html;charset=utf-8\r\n"
	responseStr += "Content-Encoding:gzip\r\n"
	responseStr += "Cache-Control:max-age=120\r\n"
	responseStr += "Age:79\r\n"
	//responseStr += "Transfer-Encoding:chunked\r\n"
	responseStr += "\r\n"

	responseBytes := []byte(responseStr)
	responseBytes = append(responseBytes, content...)
	_, err = client.Conn.Write(responseBytes)
	if err != nil {
		fmt.Println(err, requestStr)
	}
}
