package main

import (
	"encoding/base64"
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	ProxyAuthHeader = "Proxy-Authorization"
)

func main() {

	proxyURL := getEnvCaseInsensitive("http_proxy")
	username := getEnvCaseInsensitive("http_proxy_username")
	password := getEnvCaseInsensitive("http_proxy_password")
	port := getEnvCaseInsensitive("middle_proxy_port")
	if port == "" {
		port = "9000"
	}

	middleProxy := goproxy.NewProxyHttpServer()
	middleProxy.Verbose = true
	middleProxy.Tr.Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(proxyURL)
	}
	connectReqHandler := func(req *http.Request) {
		addBasicAuthHeader(username, password, req)
	}

	middleProxy.ConnectDial = middleProxy.NewConnectDialToProxyWithHandler(proxyURL, connectReqHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, middleProxy))
}

func addBasicAuthHeader(username, password string, req *http.Request) {
	req.Header.Set(ProxyAuthHeader, fmt.Sprintf("Basic %s", basicAuthBase64(username, password)))
}

func basicAuthBase64(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func getEnvCaseInsensitive(key string) string {
	keys := []string{strings.ToLower(key), strings.ToUpper(key)}
	for _, v := range keys {
		env := os.Getenv(v)
		if env != "" {
			return env
		}
	}
	return ""
}
