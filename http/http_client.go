package http

import (
	"time"
	"net"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"fmt"
	"net/http"
	"context"
	"errors"
)

var NoKeepAliveClient = BuildNonKeepAliveClient()

var KeepAliveClient = BuildKeepAliveClient(60*time.Second, 10)

type HttpClient struct {
	Client  *http.Client
	Timeout time.Duration
}

/* Constructors */
func BuildNonKeepAliveClient() *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableCompression:  true,
				DisableKeepAlives:   true,
				MaxIdleConnsPerHost: -1,
			},
		},
		Timeout: 5 * time.Second,
	}
}

func BuildKeepAliveClient(dialTimeout time.Duration, idleConnectionsPerHost int) *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   dialTimeout,
					KeepAlive: dialTimeout * 60,
				}).Dial,
				MaxIdleConnsPerHost: idleConnectionsPerHost,
			},
		},
		Timeout: 5 * time.Second,
	}
}

func (self *HttpClient) DoGet(url string, unmarshalledResponse interface{}) (statusCode int, err error) {
	return self.DoRequest("GET", url, nil, unmarshalledResponse)
}

/*
	Makes a Post Request with Given Url & Body under specified timeout.
	Incase of Success you will recieve unmarshalled Response or error otherwise
 */
func (self *HttpClient) DoPost(url string, body interface{}, unmarshalledResponse interface{}) (statusCode int, err error) {
	return self.DoRequest("POST", url, body, unmarshalledResponse)
}

func (self *HttpClient) DoRequest(method string, url string, body interface{}, unmarshalledResponse interface{}) (statusCode int, err error) {
	var requestBody, responseBytes []byte
	var request *http.Request
	var response *http.Response

	/* Encode Json */
	if requestBody, err = json.Marshal(body); err == nil {

		/* Build Request */
		if request, err = http.NewRequest(method, url, bytes.NewReader(requestBody)); err == nil {
			timeoutContext, cancelFunction := context.WithTimeout(context.Background(), self.Timeout)

			/* Set Content Type Header */
			request.Header.Set("Content-Type", "application/json")

			/* Execute Request */
			defer cancelFunction()
			if response, err = self.Client.Do(request.WithContext(timeoutContext)); err == nil {
				defer response.Body.Close()

				/* Check If Request was Successful */
				statusCode = response.StatusCode
				if response.StatusCode == http.StatusOK {
					/* Read Body & Decode */
					if responseBytes, err = ioutil.ReadAll(response.Body); err == nil {
						err = json.Unmarshal(responseBytes, unmarshalledResponse)
					}
				} else {
					err = errors.New(fmt.Sprintf("Non 200 Response. Status Code: %v", response.StatusCode))
				}
			}
		}
	}

	return
}
