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
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"bytes"
	"net"
	"strings"
	"mime"
	"path"
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

	//fmt.Println(ConfigData)
	hostInfo := ConfigData["webs"]
	fmt.Println(hostInfo)

	fmt.Println(ConfigData["php_cgi"])
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
			httpHandle := HttpHandle{"G:\\code\\godoc"}
			client.SetResponse(httpHandle)
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

type HttpHandle struct {
	WebRoot string
}

func (httpHandle HttpHandle) HandleMethod(request *httpserver.Request) (content []byte, err error) {
	filePath := httpHandle.GetAbsoluteFilePath(request)
	if isFileExists := fileoperator.FileExists(filePath); !isFileExists {
		content := httpHandle.NotFound()
		return content, nil
	}

	content, err = httpHandle.FileHandle(request)
	if err != nil {
		return nil, err
	}

	// maybe get by last modification time of file ?
	serverSideETag := generateETag(string(content), false)
	respHeader := httpserver.Header{}
	response := &httpserver.Response{ Proto: request.Proto, StatusCode: httpserver.StatusOK }

	if ifNoneMatch := request.Header.GetAll("If-None-Match"); len(ifNoneMatch) > 0 {
		for _, eTagToMatch := range ifNoneMatch {
			fmt.Println(fmt.Sprintf("sEtag: %s, clientETag: %s", serverSideETag, eTagToMatch))
			if strings.TrimSpace(eTagToMatch) == serverSideETag {
				response.StatusCode = httpserver.StatusNotModified
				break
			}
		}
	}

	// enable ETag or not according to config?
	respHeader.Set("ETag", serverSideETag)

	fp, err := os.OpenFile(filePath, os.O_RDONLY, 444)
	defer fp.Close()

	if err != nil {
		return nil, err
	}

	fileInfo, err := fileoperator.GetFileInfo(fp)

	if err != nil {
		return nil, err
	}

	respHeader.Set("Server", "f1yz/0.0.1")
	respHeader.Set("Last-Modified", fileInfo.ModTime().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	respHeader.Set("Expires", fileInfo.ModTime().Add(time.Duration(3600000)).Format("Mon, 02 Jan 2006 15:04:05 GMT"))


	//if httpHandle.IsGzipEnabled() {
	//	content = httpHandle.GzipEncoding(content)
	//	respHeader.Set("Content-Encoding", "gzip")
	//}

	// say we accept range request
	respHeader.Set("Accept-Ranges", "bytes")
	if request.IsRangeRequest() {
		fileSize := len(content)

		startEndRange := request.GetStartEndRange()
		expectedStart := 0
		expectedEnd := 0

		fmt.Println("YYYYY")
		fmt.Println(startEndRange)
		if len(startEndRange) == 2 {
			expectedStart = startEndRange[0]
			expectedEnd = startEndRange[1]
		} else {
			expectedStart = startEndRange[0]
			expectedEnd = fileSize
		}

		if expectedEnd < expectedStart || expectedEnd > fileSize {
			return httpHandle.RangeNotSatisfiable(), nil
		}

		content = content[expectedStart:expectedEnd]
		response.StatusCode = httpserver.StatusPartialContent
		respHeader.Set("Content-Range", fmt.Sprintf("bytes %v-%v/%v", expectedStart, expectedEnd, fileSize))
	}

	// parse mimeType
	if mimeType := mime.TypeByExtension(path.Ext(filePath)); mimeType != "" {
		// set response header.
		respHeader.Set("Content-Type", mimeType)
	} else {
		respHeader.Set("Content-Type", "application/octet-stream")
	}


	response.SetHeaders(respHeader)

	if httpserver.StatusNotModified != response.StatusCode {
		response.SetBody(content)
	}

	fmt.Println(response)

	content = httpHandle.WriteResponse(response)

	fmt.Println("YYYYssssssssssssssss")
	fmt.Println(content)
	return
}

func (httpHandle HttpHandle) WriteResponse(response *httpserver.Response) (content []byte) {
	statusLine := response.WriteStatusLine()
	headerStr := response.WriteHeaders()

	responsePartials := []string{
		statusLine,
		headerStr,
	}

	responseStr := strings.Join(responsePartials, "\r\n")
	responseStr = responseStr + "\r\n\r\n"

	responseBytes := []byte(responseStr)
	content = append(responseBytes, response.Body...)
	return
}

func (httpHandle HttpHandle) GetAbsoluteFilePath(request *httpserver.Request) string {
	return httpHandle.WebRoot + request.RequestURI
}

func (httpHandle HttpHandle) FileHandle(request *httpserver.Request) (content []byte, err error) {
	fullPath := httpHandle.GetAbsoluteFilePath(request)
	content, err = fileoperator.ReadAll(fullPath)

	return
}

func (httpHandle HttpHandle) CgiHandle(request *httpserver.Request) (content []byte, err error) {
	cgiConn, err := net.Dial("tcp", "127.0.0.1:9001")
	cgiConn.SetReadDeadline(time.Now().Add(3 * time.Second))
	cgiConn.Write([]byte("123"))

	cgiContent := make([]byte, 512)
	var n int
	for {
		n, err = cgiConn.Read(cgiContent)
		if err != nil || n <= 0 {
			cgiConn.SetReadDeadline(time.Time{})
			break
		}

		content = append(content, cgiContent...)
	}

	return
}

func (httpHandle HttpHandle) GzipEncoding(content []byte) (gzipContent []byte) {
	var b bytes.Buffer
	gzipWriter := gzip.NewWriter(&b)
	defer gzipWriter.Close()

	gzipWriter.Write(content)
	gzipWriter.Flush()

	gzipContent = b.Bytes()

	return
}

func (httpHandle HttpHandle) GenerateETag(requestURI string) (string, error) {
	fullPath := httpHandle.WebRoot + requestURI

	fp, err := os.OpenFile(fullPath, os.O_RDONLY, 444)
	defer fp.Close()

	if err != nil {
		return "", err
	}

	fileInfo, err := fileoperator.GetFileInfo(fp)
	if err != nil {
		return "", err
	}

	fileLastModifiedAt := fileInfo.ModTime().Unix()
	eTag := generateETag(string(fileLastModifiedAt), false)
	return eTag, nil
}

func (httpHandle *HttpHandle) NotFound() []byte {

	notFoundStr := "HTTP/1.1 404 NOT FOUND"
	return []byte(notFoundStr)
}

func (httpHandle *HttpHandle) RangeNotSatisfiable() []byte {
	return []byte("HTTP/1.1 416 Requested Range Not Satisfiable")
}

func (httpHandle *HttpHandle) IsGzipEnabled() bool {
	return true
}


func generateETag(identityStr string, weak bool) string {
	var eTag string

	md5 := md5.New()
	md5.Write([]byte(identityStr))
	eTag = hex.EncodeToString(md5.Sum(nil))

	if weak {
		eTag = "W/" + eTag
	}

	return eTag
}
