# Example of a secure web application with Gin

This is an example of a secure web application with Gin. It includes the following security headers:

- Content-Security-Policy
- Permissions-Policy
- Referrer-Policy
- Strict-Transport-Security
- X-Frame-Options
- X-Xss-Protection
- X-Content-Type-Options

Also the web application has strict Host Header to avoid SSRF and Host Header Injection.

1. Security Headers Example

```
christofherkost at L-EEM091 in ~/D/s/c/gin-examples
↪ curl http://localhost:8080 -I
HTTP/1.1 404 Not Found
Content-Security-Policy: default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';
Content-Type: text/plain
Permissions-Policy: geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()
Referrer-Policy: strict-origin
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-Xss-Protection: 1; mode=block
Date: Thu, 28 Mar 2024 11:38:05 GMT
Content-Length: 18
```

2. Host Header Injection Example
```
christofherkost at L-EEM091 in ~/D/s/c/gin-examples
↪ curl http://localhost:8080 -I -H "Host:neti.ee"
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Thu, 28 Mar 2024 11:38:23 GMT
Content-Length: 31
```