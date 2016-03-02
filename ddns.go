package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jmoiron/jsonq"
	"github.com/xwjdsh/httphelper"
)

type DDNS struct {
	Config *Config
	http   *httpHelper.HttpHelper
}

func (this *DDNS) initHttp() {
	this.http.Host = "https://dnsapi.cn/"
	this.http.CommonHeader.Add("UserAgent", fmt.Sprintf("ddns/0.1 (%s)", this.Config.Email))
	params := this.http.CommonParam

	if this.Config.Token != "" {
		params.Add("login_token", this.Config.Token)
	} else {
		params.Add("login_email", this.Config.Email)
		params.Add("login_password", this.Config.Password)
	}
	params.Add("format", "json")
	params.Add("lang", "cn")
	params.Add("error_on_empty", "no")
}

func (this *DDNS) domainID() (int, error) {
	param := url.Values{}
	param.Add("type", "all")
	param.Add("offset", "0")
	param.Add("length", "20")
	param.Add("keyword", this.Config.Domain)
	resp, err := this.http.Send("POSTFORM", "Domain.List", param, nil)
	if err != nil {
		return 0, errors.New("获取Domain.List错误，请检查!")
	}
	data := map[string]interface{}{}
	dec := json.NewDecoder(bytes.NewReader(resp))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	if code, _ := jq.Int("status", "code"); code != 1 {
		message, _ := jq.String("status", "message")
		return 0, errors.New(message)
	}
	id, err := jq.Int("domains", "0", "id")
	if err != nil {
		return 0, errors.New("没有指定的域名，请检查!")
	}
	return id, nil
}

func (this *DDNS) recordID(domainID int) (int, string, error) {
	param := url.Values{}
	param.Add("domain_id", strconv.Itoa(domainID))
	param.Add("sub_domain", this.Config.SubDomain)
	resp, err := this.http.Send("POSTFORM", "Record.List", param, nil)
	if err != nil {
		return 0, "", errors.New("获取Record.List错误，请检查!")
	}
	data := map[string]interface{}{}
	dec := json.NewDecoder(bytes.NewReader(resp))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	if code, _ := jq.Int("status", "code"); code != 1 {
		message, _ := jq.String("status", "message")
		return 0, "", errors.New(message)
	}
	id, err := jq.Int("records", "0", "id")
	value, err := jq.String("records", "0", "value")
	if err != nil {
		return 0, "", errors.New("没有指定的域名记录，请检查!")
	}
	return id, value, nil
}
