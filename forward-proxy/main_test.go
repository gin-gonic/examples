package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestReverse(t *testing.T) {
	for _, spoofed := range []bool{false, true} {
		name := "request and response"
		if spoofed {
			name = "untrusted forwarding and hop headers"
		}
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequestWithContext(
				t.Context(),
				http.MethodPost,
				"http://proxy.example/items/a%2Fb?q=one&q=two",
				strings.NewReader("payload"),
			)
			req.RemoteAddr = "192.0.2.10:1234"
			req.Header.Set("X-Custom", "kept")
			if spoofed {
				req.Header.Set("Forwarded", "for=attacker")
				req.Header.Set("X-Forwarded-For", "attacker")
				req.Header.Set("X-Forwarded-Host", "attacker.example")
				req.Header.Set("X-Forwarded-Proto", "https")
				req.Header.Set("Connection", "X-Hop")
				req.Header.Set("X-Hop", "remove")
			}
			originalHeaders := req.Header.Clone()
			originalURL := req.URL.String()
			originalTransport := http.DefaultTransport
			t.Cleanup(func() { http.DefaultTransport = originalTransport })
			called := false
			http.DefaultTransport = roundTripFunc(func(out *http.Request) (*http.Response, error) {
				called = true
				if out.URL.String() != "http://xxx.xxx.xxx/items/a%2Fb?q=one&q=two" ||
					(out.Host != "" && out.Host != "xxx.xxx.xxx") ||
					out.Method != http.MethodPost {
					t.Errorf(
						"unexpected outbound request: %s %s host=%s",
						out.Method,
						out.URL,
						out.Host,
					)
				}
				body, err := io.ReadAll(out.Body)
				if err != nil || string(body) != "payload" {
					t.Errorf("body=%q err=%v", body, err)
				}
				want := map[string]string{
					"X-Custom":          "kept",
					"X-Forwarded-For":   "192.0.2.10",
					"X-Forwarded-Host":  "proxy.example",
					"X-Forwarded-Proto": "http",
					"Forwarded":         "",
					"Connection":        "",
					"X-Hop":             "",
				}
				for k, v := range want {
					if got := out.Header.Get(k); got != v {
						t.Errorf("%s=%q want %q", k, got, v)
					}
				}
				return &http.Response{
					StatusCode: http.StatusCreated,
					Header:     http.Header{"X-Upstream": {"ok"}},
					Body:       io.NopCloser(strings.NewReader("upstream response")),
				}, nil
			})
			router := gin.New()
			router.Any("/*proxyPath", Reverse)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			if !called {
				t.Fatal("upstream transport was not called")
			}
			if rec.Code != http.StatusCreated || rec.Body.String() != "upstream response" ||
				rec.Header().Get("X-Upstream") != "ok" {
				t.Errorf(
					"unexpected response: code=%d body=%q headers=%v",
					rec.Code,
					rec.Body.String(),
					rec.Header(),
				)
			}
			if req.URL.String() != originalURL || req.Host != "proxy.example" ||
				!reflect.DeepEqual(req.Header, originalHeaders) {
				t.Error("reverse proxy modified the inbound URL, host or headers")
			}
		})
	}
}
