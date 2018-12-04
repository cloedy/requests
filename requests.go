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
