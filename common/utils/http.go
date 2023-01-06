package utils

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// http请求方法
const (
	POST = "POST"
	GET  = "GET"
	PUT  = "PUT"
	DEL  = "DELETE"
	HEAD = "HEAD"
)

const (
	httpRequestTimeOut = 3 * time.Second // 回调接口主动推送响应超时时间
)

// Get 常规GET请求
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if resp != nil {
		log.Println(resp.StatusCode, url)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Post 上传
func Post(url, payload string) ([]byte, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(payload)) //"name=cjb"
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// PostData 表单模式传输数据
func PostData(url string, payload url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// PostFile 表单模式传输文件
func PostFile(method, url, path, paramName string, params map[string]string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reqbody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqbody)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	resp, err := http.NewRequest(method, url, reqbody)
	resp.Header.Set("Content-Type", writer.FormDataContentType())
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

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
