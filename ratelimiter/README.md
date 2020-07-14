# Ratelimit in Gin

This project is a sample for ratelimit using Leaky Bucket. Although the golang official pkg provide a implement with Token Bucket [time/rate](https://pkg.go.dev/golang.org/x/time/rate?tab=doc),

you can also make your owns via gin's functional `Use()` to integrate extra middlewares.

## Effect

```go
// You can assign the ratelimit of the server
// rps: requests per second
go run rate.go -rps=100
```

- Let's hava a simple test by ab with 3000 mock requests, not surprisingly，it will takes 10ms each request.

```bash
ab -n 3000 -c 1 http://localhost:8080/rate
```

- Gin Log Output

```bash
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    8.307734ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |   10.512913ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |     8.54681ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    8.356436ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    9.677276ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    7.536156ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    11.57084ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |       7.802ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    9.602394ms |       127.0.0.1 | GET      /rate
```

- AB Test Reporter

```java
Concurrency Level:      1
Time taken for tests:   30.00 seconds
Complete requests:      3000
Requests per second:    100.00 [#/sec] (mean)
Time per request:       10.001 [ms] (mean)
Time per request:       10.001 [ms] (mean, across all concurrent requests)
```
