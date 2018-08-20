package cloedy

import (
	"net/http"
	"strings"
	"bytes"
	"fmt"
	"io/ioutil"
)

type RequestParams struct {
	Method, Url, Data string
	Headers           map[string]string
	Files             []bytes.Reader
}

func Requests(params RequestParams) (*http.Response, error) {
	method := params.Method
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

func Text(response *http.Response) string {
	t, _ := ioutil.ReadAll(response.Body)
	return fmt.Sprintf("%s", t)
}
