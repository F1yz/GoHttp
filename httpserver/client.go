package httpserver

import (
	"net"
	"fmt"
	"io"
)

type Client struct {
	Conn net.Conn
	WaitChan chan string
}

type ParseRequest interface {
	parseRequest(requestStr string) map[string]string
}

func (client *Client) GetRequest() {
	requestByte := make([]byte, 512)
	requestStr := ""
	contentLength := 0;

	for {
		len, err := client.Conn.Read(requestByte)
		fmt.Println("len", len, err, client.Conn.RemoteAddr())
		if err != nil {
			if err == io.EOF  {
				break
			}
			fmt.Println("err:",err)
		}

		contentLength += len
		requestStr += string(requestByte)

		if len < 512 {
			break
		}
	}

	client.WaitChan <- requestStr
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

	fmt.Println([]byte(responseStr))
	len, err := client.Conn.Write([]byte(responseStr))
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(len, requestStr)
}
