package cloedy

import (
	"net/http"
	"strings"
	"io"
)

type RequestParams struct {
	Method, Url, Data string
	Headers           map[string]string
	Files             io.Reader
}

var Header = map[string]string{
	"Connection":                "keep-alive",
	"Pragma":                    "no-cache",
	"Cache-Control":             "no-cache",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
	"Accept-Encoding":           "gzip, deflate",
	"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8",
}

func method(m string) string {
	switch strings.ToLower(m) {
	case "get":
		return http.MethodGet
	case "post":
		return http.MethodPost
	case "options":
		return http.MethodOptions
	case "delete":
		return http.MethodDelete
	default:
		return http.MethodGet
	}
}

func (params *RequestParams) Requests() (*http.Response, error) {
	method := method(params.Method)
	url := params.Url
	data := params.Data
	datas := strings.NewReader(data)
	req, err := http.NewRequest(method, url, datas)
	if err != nil {
		return new(http.Response), err
	}
	if params.Headers != nil {
		for name, value := range params.Headers {
			req.Header.Add(name, value)
		}
	}
	client := http.DefaultClient
	return client.Do(req)
}
