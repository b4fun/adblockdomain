package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func tryUrl(rurl string) (io.ReadCloser, error) {
	_, err := url.Parse(rurl)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(rurl)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func tryFile(path string) (io.ReadCloser, error) {
	return os.Open(path)
}
