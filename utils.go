package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func currentIP() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", errors.New("获取当前公网发生IP错误!")
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("解析响应出错!")
	}
	return strings.TrimSpace(string(result[:])), nil
}
