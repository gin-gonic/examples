package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(ForwardMid)

	// Create a catchall route
	r.Any("/*proxyPath", Reverse)

	if err := r.Run(":8888"); err != nil {
		panic(err)
	}
}

func ForwardMid(c *gin.Context) {
	// !!! adapt to your request header set
	if v, ok := c.Request.Header["Forward"]; ok {
		if v[0] == "ok" {
			resp, err := http.DefaultTransport.RoundTrip(c.Request)
			if err != nil {
				http.Error(c.Writer, err.Error(), http.StatusServiceUnavailable)
				c.Abort()
				return
			}
			defer resp.Body.Close()
			copyHeader(c.Writer.Header(), resp.Header)
			c.Writer.WriteHeader(resp.StatusCode)
			io.Copy(c.Writer, resp.Body)
			c.Abort()
			return
		}
	}

	c.Next()
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func Reverse(c *gin.Context) {
	remote, _ := url.Parse("http://xxx.xxx.xxx")
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Host = remote.Host
		req.URL.Scheme = remote.Scheme
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
