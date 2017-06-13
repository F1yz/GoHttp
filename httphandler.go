package main

import (
	"fmt"
	"httpserver"
)

func main() {
	server := httpserver.StartServer("127.0.0.1", 8890)
	err := server.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	for  {
		client, err := server.GetClient()
		if (err != nil) {
			fmt.Println(err, 123)
		}

		go func() {
			client.GetRequest()
		}()

		go func() {
			client.SetReponse()
		}()
	}
}
