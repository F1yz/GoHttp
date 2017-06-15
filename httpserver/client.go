package httpserver

import (
	"net"
	"fmt"
	"io"
)

type Client struct {
	Conn net.Conn
	WaitChan chan map[string]string
}

type ParseRequest interface {
	ParseRequest(requestBytes []byte) map[string]string
}

func (client *Client) GetRequest(request ParseRequest) {
	var requestBytes []byte
	requestByte := make([]byte, 512)

	for {
		len, err := client.Conn.Read(requestByte)
		if err != nil {
			if err == io.EOF  {
				break
			}
			fmt.Println(err)
		}

		if len < 512 {
			requestBytes = append(requestBytes, requestByte[:len]...)
			break
		}
		requestBytes = append(requestBytes, requestByte...)
	}

	requestData := request.ParseRequest(requestBytes)
	client.WaitChan <- requestData
}

func (client *Client) SetReponse() {
	defer func() {
		client.Conn.Close()
		close(client.WaitChan)
	}()

	requestStr := <-client.WaitChan

	responseStr := "HTTP/1.1 200 OK\r\n"
	responseStr += "Expires:Tue, 13 Jun 2017 11:57:00 GMT\r\n"
	responseStr += "Content-Type:text/html;charset=utf-8\r\n"
	responseStr += "Cache-Control:max-age=120\r\n"
	responseStr += "Age:79\r\n"
	//responseStr += "Transfer-Encoding:chunked\r\n"
	responseStr += "\r\n"
	responseStr += "你好\r\n"

	_, err := client.Conn.Write([]byte(responseStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(requestStr)
}
