// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "gae-goa-sample": hello Resource Client
//
// Command:
// $ goagen
// --design=github.com/y-ogura/gae-goa-sample/design
// --out=$(GOPATH)/src/github.com/y-ogura/gae-goa-sample
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// HelloWorldHelloPath computes a request path to the HelloWorld action of hello.
func HelloWorldHelloPath() string {

	return fmt.Sprintf("/hello")
}

// HelloWorld
func (c *Client) HelloWorldHello(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewHelloWorldHelloRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewHelloWorldHelloRequest create the request corresponding to the HelloWorld action endpoint of the hello resource.
func (c *Client) NewHelloWorldHelloRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
