package main

import (
	"errors"
	"net/url"

	"github.com/xwjdsh/httphelper"
)

type Config struct {
	Email     string `json:"login_email"`
	Password  string `json:"login_password"`
	Token     string `json:"login_token"`
	Domain    string `json:"domain"`
	SubDomain string `json:"sub_domain"`

	checked bool
}

func (this *Config) newDDNS() (*DDNS, error) {
	if !(this.Domain != "" && ((this.Email != "" && this.Password != "") || this.Token != "")) {
		return nil, errors.New("Incorrect config!")
	}
	this.checked = true
	helper := &httpHelper.HttpHelper{CommonHeader: url.Values{}, CommonParam: url.Values{}, Log: true}
	return &DDNS{
		Config: this,
		http:   helper,
	}, nil
}
