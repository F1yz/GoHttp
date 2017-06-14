package main

import (
	"fmt"
	"httpserver"
	"configloader"
	"os"
	"fileoperator"
	"errors"
)

var ConfigData map[interface{}]interface{}
var ConfigLoader ConfigureLoader

type ConfigureLoader interface {
	LoadConfigure(configData []byte) (map[interface{}]interface{}, error)
	SetConfigure(key interface{}, setConfigData interface{}) error
}

func main() {
	setConfigureLoader()
	configBytes, errMsg := getConfigBytes();
	if errMsg != nil {
		fmt.Println(errMsg)
		os.Exit(1)
	}

	errMsg = LoadConfigure(configBytes)
	if errMsg != nil {
		fmt.Println(errMsg)
		os.Exit(2)
	}

	server := httpserver.StartServer(ConfigData["address"].(string), int(ConfigData["port"].(int)))
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

func getConfigBytes() (readBytes[]byte, err error) {
	configPath, err := getConfigPath()
	if err != nil {
		return
	}

	readBytes, err = readConfig(configPath)
	return
}

func getConfigPath () (configPath string, err error) {
	configPath = os.Getenv("ONLYFUNCONFIG")
	if configPath == "" {
		err = errors.New("请设置配置文件环境变量(ONLYFUNCONFIG)");
	}
	return
}

func readConfig(configPath string) (readBytes[]byte, err error) {
	readBytes, err = fileoperator.ReadAll(configPath)
	return
}

func setConfigureLoader() {
	ConfigLoader = &configloader.YamlLoader{}
}

func LoadConfigure(configBytes []byte) (err error) {
	ConfigData, err = ConfigLoader.LoadConfigure(configBytes)
	return
}

func SetConfigure(key interface{}, setConfigData interface{}) (err error) {
	err = ConfigLoader.SetConfigure(key, setConfigData)
	return
}