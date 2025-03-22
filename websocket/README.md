# WebSocket Example

This project demonstrates a simple WebSocket implementation using the Gin framework and Gorilla WebSocket library in Go. It includes a server that handles WebSocket connections and a client that connects to the server and sends periodic messages.

## Project Overview

The project consists of two main components:

- **Server**: A WebSocket server that echoes messages back to the client.
- **Client**: A WebSocket client that connects to the server, sends messages, and handles responses.

## Setup Instructions

1. **Clone the repository**:

   ```sh
   git clone https://github.com/your-repo/websocket-example.git
   cd websocket-example/websocket
   ```

2. **Install dependencies**:
   Ensure you have Go installed, then run:

   ```sh
   go mod tidy
   ```

3. **Run the server**:

   ```sh
   go run server/server.go
   ```

4. **Run the client**:
   Open a new terminal and run:

   ```sh
   go run client/client.go
   ```

## Usage Instructions

### Server

The server listens on port 8080 by default and provides two endpoints:

- `/echo`: Handles WebSocket connections and echoes messages back to the client.
- `/`: Serves an HTML page for testing the WebSocket connection.

### Client

The client connects to the server's `/echo` endpoint and sends periodic messages. It also handles incoming messages and logs them to the console.

## Code Explanation

### Server Implementation

The server code is located in `server/server.go`. It uses the Gin framework to set up HTTP routes and the Gorilla WebSocket library to handle WebSocket connections.

- **`echo` function**: Upgrades HTTP connections to WebSocket and echoes received messages back to the client.
- **`home` function**: Serves an HTML page for testing the WebSocket connection.
- **`main` function**: Sets up the Gin router, defines routes, and starts the server.

### Client Implementation

The client code is located in `client/client.go`. It connects to the server's WebSocket endpoint and sends periodic messages.

- **`main` function**: Parses command-line flags, sets up signal handling, and establishes a WebSocket connection. It sends periodic messages and handles incoming messages.
