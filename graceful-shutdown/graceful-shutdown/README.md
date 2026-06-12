# Graceful Shutdown Examples

This directory contains examples demonstrating how to implement graceful shutdowns in a Gin server using context with and without context.

## Project Structure

- `notify-with-context/`: Example of graceful shutdown using context.
- `notify-without-context/`: Example of graceful shutdown without using context.

## Usage

### Notify with Context

1. Install the required dependencies:

   ```bash
   go get -u github.com/gin-gonic/gin
   ```

2. Run the server:

   ```bash
   go run notify-with-context/server.go
   ```

3. Access the server at `http://localhost:8080/`.

4. To trigger a graceful shutdown, send an interrupt signal (e.g., `Ctrl+C` in the terminal). The server will complete any ongoing requests before shutting down.

### Notify without Context

1. Install the required dependencies:

   ```bash
   go get -u github.com/gin-gonic/gin
   ```

2. Run the server:

   ```bash
   go run notify-without-context/server.go
   ```

3. Access the server at `http://localhost:8080/`.

4. To trigger a graceful shutdown, send an interrupt signal (e.g., `Ctrl+C` in the terminal). The server will complete any ongoing requests before shutting down.

## Code Explanation

### Notify with Context Example

- The server is initialized with a simple route that simulates a delay of 10 seconds.
- A context is created that listens for interrupt signals.
- When an interrupt signal is received, the server is shut down gracefully using the context.

### Notify without Context Example

- The server is initialized with a simple route that simulates a delay of 5 seconds.
- A channel is created to listen for interrupt signals.
- When an interrupt signal is received, the server is shut down gracefully using the `server.Shutdown()` method with a timeout context.
