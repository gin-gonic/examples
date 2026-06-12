# Graceful Shutdown Example - Close Method

This example demonstrates how to implement a graceful shutdown in a Gin server using the `server.Close()` method.

## Project Structure

- `server.go`: The main server implementation that handles graceful shutdown.

## Usage

1. Install the required dependencies:

   ```bash
   go get -u github.com/gin-gonic/gin
   ```

2. Run the server:

   ```bash
   go run server.go
   ```

3. Access the server at `http://localhost:8080/`.

4. To trigger a graceful shutdown, send an interrupt signal (e.g., `Ctrl+C` in the terminal). The server will complete any ongoing requests before shutting down.

## Code Explanation

- The server is initialized with a simple route that simulates a delay of 5 seconds.
- A channel is created to listen for interrupt signals.
- When an interrupt signal is received, the server is closed gracefully using the `server.Close()` method.
