package main

import (
	"fmt"
	"os"
	"errors"
	"strconv"
	"runtime"
	"configloader"
	"time"
	"fileoperator"
	"httpserver"
	"httpparse"
)

var ConfigData map[interface{}]interface{}
var configLoader configloader.ConfigureLoader
var parser *httpparse.HttpParse

func main() {
	configBytes, errMsg := getConfigBytes();
	if errMsg != nil {
		fmt.Println(errMsg)
		os.Exit(1)
	}

	errMsg = loadConfigure(configBytes)
	if errMsg != nil {
		fmt.Println(errMsg)
		os.Exit(2)
	}

	errMsg = savePidFile()
	if errMsg != nil {
		fmt.Println(errMsg)
		os.Exit(3)
	}

	setParser()
	setProcsNum()

	server := httpserver.StartServer(ConfigData["address"].(string), int(ConfigData["port"].(int)))
	err := server.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	for  {
		client, err := server.GetClient(ConfigData["request_read_buffer"].(int), time.Duration(ConfigData["life_time"].(int)))
		if (err != nil) {
			fmt.Println(err)
		}

		go func() {
			client.GetRequest(parser)
		}()

		go func() {
			client.SetReponse()
		}()
	}
}

func savePidFile() (err error) {
	pid := getPid()
	filePath := ConfigData["pidfile"].(string)
	err = savePid(filePath, pid)
	return
}


func getPid () (pid int) {
	pid = os.Getpid()
	return
}

func savePid(filePath string, pid int) (err error) {
	err = fileoperator.WriteIn(filePath, strconv.Itoa(pid))
	return
}

func getConfigBytes() (readBytes[]byte, err error) {
	setConfigureLoader()
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
	configLoader = &configloader.YamlLoader{}
}

func loadConfigure(configBytes []byte) (err error) {
	ConfigData, err = configLoader.LoadConfigure(configBytes)
	return
}

func setParser() {
	parser = &httpparse.HttpParse{}
}

func setProcsNum() {
	procsNum := ConfigData["procss"].(int)
	if procsNum == 0 {
		procsNum = runtime.NumCPU() / 2
	}

	runtime.GOMAXPROCS(procsNum)
}

func SetConfigure(key interface{}, setConfigData interface{}) (err error) {
	err = configLoader.SetConfigure(key, setConfigData)
	return
}