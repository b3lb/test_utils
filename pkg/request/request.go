package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func initRequestContext(c *gin.Context) {
	if c.Request == nil {
		c.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
	}
}

// AddRequestBody create a request body from a json string
func AddRequestBody(c *gin.Context, data string) {
	initRequestContext(c)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(data)))
}

// SetRequestMethod set the request method
func SetRequestMethod(c *gin.Context, method string) {
	initRequestContext(c)
	c.Request.Method = method
}

// SetRequestContentType set the request content type
func SetRequestContentType(c *gin.Context, contentType string) {
	initRequestContext(c)
	c.Request.Header.Set("Content-Type", contentType)
}

// SetRequestParams set the request url
func SetRequestParams(c *gin.Context, value string) {
	initRequestContext(c)
	c.Request.URL = &url.URL{
		RawQuery: value,
	}
}

// SetRequestHeader set the request header
func SetRequestHeader(c *gin.Context, key, value string) {
	initRequestContext(c)
	c.Request.Header.Set(key, value)
}

// SetRequestHost set the request host
func SetRequestHost(c *gin.Context, hostname string) {
	initRequestContext(c)
	c.Request.Host = hostname
}
