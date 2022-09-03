# Gin websocket Client and server example

This example shows a simple client and server.

The server echoes messages sent to it. The client sends a message every second and prints all messages received.

To run the example, start the server:

```bash
go run server/server.go
```

Next, start the client:

```bash
go run client/client.go
```

The server includes a simple web client. To use the client, open [URL](http://127.0.0.1:8080) in the browser and follow the instructions on the page.

## deps

- [gorilla/websocket](https://github.com/gorilla/websocket)
