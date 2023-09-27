#go-socket

A simple WebSocket server built with GoLang and Gin, integrated with RabbitMQ message broker for real-time messaging.

## Features

- WebSocket server sending periodic messages to clients.
- Integration with RabbitMQ for broadcasting messages to connected clients.
- Graceful shutdown mechanism.
- Environment variable configuration.

## Prerequisites

- Go 1.15 or higher
- RabbitMQ server
- [Gin](https://github.com/gin-gonic/gin) and [Gorilla WebSocket](https://github.com/gorilla/websocket) dependencies

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/ahmedMHasan/go-socket.git
   ```

2. Navigate to the project directory:

   ```shell
   cd go-socket
   ```

3. Set up your environment variables:

   ```shell
   export AMQP_URL=your_rabbitmq_url
   ```

4. Build and run the server:

   ```shell
   go run main.go
   ```

5. Access the WebSocket at `ws://localhost:8080/ws` from your WebSocket client.

## Configuration

You can configure the following environment variables:

- `AMQP_URL`: RabbitMQ server URL.

## Usage

### WebSocket Client Example

```javascript
const socket = new WebSocket('ws://localhost:8080/ws');

socket.addEventListener('open', (event) => {
  console.log('Connected to the WebSocket server');
});

socket.addEventListener('message', (event) => {
  console.log('Received message:', event.data);
});

socket.addEventListener('close', (event) => {
  console.log('WebSocket connection closed:', event.reason);
});

socket.addEventListener('error', (event) => {
  console.error('WebSocket error:', event.error);
});
```

## Graceful Shutdown

To gracefully stop the server, send a termination signal (e.g., `SIGINT` or `SIGTERM`) to the running process. The server will handle the shutdown and clean up resources.

To send a shutdown signal in Unix-like systems:

```shell
kill -SIGINT <process_id>
```

Replace `<process_id>` with the actual process ID of the running server.

## References

- [Building a Production-Grade WebSocket for Notifications with Golang and Gin](https://medium.com/@abhishekranjandev/building-a-production-grade-websocket-for-notifications-with-golang-and-gin-a-detailed-guide-5b676dcfbd5a) by [Abhishek Ranjan]

  This Medium article provides a detailed guide on building a WebSocket server for notifications using Golang and Gin. It served as a valuable reference for implementing real-time messaging in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

