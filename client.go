package transmission

// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	// safe
	address  string
	username string
	password string
	client   *http.Client
	// unsafe
	session  string
	lastseen int64
	mu       sync.RWMutex
}

func (self *Client) GetClientAddress() string {
	return self.address
}
func (self *Client) GetClientUsername() string {
	return self.username
}
func (self *Client) GetClientPassword() string {
	return self.password
}
func (self *Client) GetClientHTTPClient() *http.Client {
	return self.client
}
func (self *Client) GetClientSession() string {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.session
}
func (self *Client) GetClientLastSeen() int64 {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.lastseen
}

func ConstructAddress(
	secure bool,
	host string,
	port int,
) string {
	if host == "" {
		host = "127.0.0.1"
	}
	if port < 1 {
		port = 9091
	}
	if secure {
		return fmt.Sprintf("https://%s:%d/transmission/rpc", host, port)
	}
	return fmt.Sprintf("http://%s:%d/transmission/rpc", host, port)
}
func Construct(
	client *http.Client,
	address string,
	username string,
	password string,
) *Client {
	if address == "" {
		address = ConstructAddress(false, "", 0)
	}
	if client == nil {
		client = createClient()
	}
	return &Client{
		address:  address,
		username: username,
		password: password,
		client:   client,
	}
}
func (self *Client) Request(
	req *Request,
) (
	*Response,
	error,
) {
	return self.request(true, req)
}
func (self *Client) request(
	parent bool,
	req *Request,
) (
	*Response,
	error,
) {
	if req == nil {
		return nil, fmt.Errorf("Transmission.Client.Request(): request nil")
	}
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Transmission.Client.Request(): failed to marshal request: %s", err)
	}
	request, err := http.NewRequest("POST", self.address, bytes.NewBuffer(bs))
	if err != nil {
		return nil, fmt.Errorf("Transmission.Client.Request(): failed to create new request: %s", err)
	}
	request.Header.Set("X-Transmission-Session-Id", self.session)
	if self.username != "" {
		request.SetBasicAuth(self.username, self.password)
	}
	res, err := self.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Transmission.Client.Request(): failed to do request: %s", err)
	}
	if res.StatusCode == 401 {
		// username or password was incorrect
		err := fmt.Errorf("Transmission.Client.Request(): Unauthorized Access!")
		log.Println(err)
		return nil, err
	}
	if res.StatusCode == 409 {
		log.Println("Transmission.Client.Request(): bad session, retrying")
		s := res.Header.Get("X-Transmission-Session-Id")
		if s == "" {
			return nil, fmt.Errorf("Transmission.Client.Request(): failed to get session")
		}
		self.mu.Lock()
		self.session = s
		self.mu.Unlock()
		log.Printf("Transmission.Client.Request(): session: %s\n", s)
		if parent {
			// this is the best way to stop runaway recursion
			return self.request(false, req)
		}
		// whatever is calling Request() should retry the attempt
		// we don't want to endlessly keep making failing requests
		return nil, fmt.Errorf("Transmission.Client.Request(): failed to make request")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Transmission.Request(): failed to read result: %s", err)
	}
	result := &Response{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Transmission.Client.Request(): failed to unmarshal result: %s", err)
	}
	self.mu.Lock()
	self.lastseen = time.Now().Unix()
	self.mu.Unlock()
	return result, nil
}
func createClient() *http.Client {
	transport := &http.Transport{}
	transport.TLSClientConfig = &tls.Config{}
	transport.TLSClientConfig.InsecureSkipVerify = true
	transport.TLSHandshakeTimeout = time.Second * 10
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
}
