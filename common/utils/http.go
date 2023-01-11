package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"stash-scrapers/common/config"
	"time"

	"golang.org/x/net/proxy"
)

// http请求方法
const (
	POST      = "POST"
	GET       = "GET"
	PUT       = "PUT"
	DEL       = "DELETE"
	HEAD      = "HEAD"
	SocksMode = "socks"
	HTTPMode  = "http"
)

const (
	httpRequestTimeOut = 30 * time.Second // 回调接口主动推送响应超时时间
	userAgentDefault   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
)

// HTTPStd 标准自定义请求,可请求二进制流，禁止网址转跳
func HTTPStd(method, url string, payload []byte) (int, []byte, error) {
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 1 {
			return errors.New("拦截到转跳请求，请检查并正确设置回调地址。")
		}
		return nil
	}
	client.Timeout = httpRequestTimeOut
	reqbody := bytes.NewReader(payload)
	req, err := http.NewRequest(method, url, reqbody)
	if err != nil {
		return -1, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	// req.Header.Set("Cookie", "name=anny")
	resp, err := client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}
	return resp.StatusCode, body, nil
}

func HTTPForMinnanoAV(url string, referer, userAgent string) (int, []byte, int, error) {
	return HTTPCheckJump(GET, url, nil, "www.minnano-av.com", referer, "")
}

// HTTPCheckJump
func HTTPCheckJump(method, urlPath string, payload []byte, host, referer, userAgent string) (int, []byte, int, error) {
	client := &http.Client{}
	var jumpNum int = 0
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		jumpNum = len(via)
		return nil
	}
	proxySet(client)
	client.Timeout = httpRequestTimeOut
	reqbody := bytes.NewReader(payload)
	req, err := http.NewRequest(method, urlPath, reqbody)
	if err != nil {
		return -1, nil, jumpNum, err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Host", host)
	if len(referer) != 0 {
		req.Header.Set("referer", referer)
	}
	if len(userAgent) == 0 {
		req.Header.Set("User-Agent", userAgentDefault)
	} else {
		req.Header.Set("User-Agent", userAgent)
	}
	req.Header.Set("Cookie", "PHPSESSID=f3lrusc05ug7jdhfkk0bp047f1;")
	resp, err := client.Do(req)
	if err != nil {
		return -1, nil, jumpNum, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, jumpNum, err
	}
	return resp.StatusCode, body, jumpNum, nil
}

func proxySet(client *http.Client) {
	if config.ProxyInfo.Type == SocksMode {
		dialer, err := proxy.SOCKS5("tcp", config.ProxyInfo.Socks, nil, proxy.Direct)
		if err != nil {
			log.Println("invalid socks proxy:", err)
			return
		}
		client.Transport = &http.Transport{
			Dial: dialer.Dial,
		}
	} else if config.ProxyInfo.Type == HTTPMode {
		proxyAddress, err := url.Parse(config.ProxyInfo.HTTP)
		if err != nil {
			log.Println("invalid http proxy:", err)
			return
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyAddress),
		}
	}
}

// HTTP 自定义请求,可请求二进制流
func HTTP(method, url string, payload []byte) ([]byte, error) {
	statusCode, body, err := HTTPStd(method, url, payload)
	if err != nil {
		return nil, err
	}
	if statusCode >= 300 && statusCode != 404 {
		log.Println(string(body))
	}
	return body, nil
}

// HTTPWithoutBody 自定义请求,可请求二进制流,不获取响应报文
func HTTPWithoutBody(method, url string, payload []byte) (int, error) {
	client := &http.Client{}
	reqbody := bytes.NewReader(payload)
	req, err := http.NewRequest(method, url, reqbody)
	if err != nil {
		return -1, err
	}
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
