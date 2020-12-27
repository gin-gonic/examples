# A simple reverse proxy

We can see real server in real_server.go and proxy server in reverse_server.go

Run this two file and if we do some request like `curl 'http://localhost:2002/something`

we will get a response in JSON contains ip of client and path we requested.
