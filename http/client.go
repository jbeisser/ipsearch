package http

import (
	"fmt"
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	client := http.Client{
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 3 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
	return &client
}

func Get(url string) (resp *http.Response, err error) {
	client := NewHttpClient()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("failed: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("req failed: %v", err)
	}
	return resp, err
}
