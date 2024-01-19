<div align="center">
  <img src="https://github.com/hmartiins/chat/assets/51277667/fca09cd1-d66a-4843-a6f4-cbbc1bfd738f" width="200" />

  <h1>Gotter Chat</h1>
  <p>💬 Application for exchanging messages in real time 📱</p>
</div>

## Project Overview

This project is a backend implementation in Go for a real-time chat system using WebSocket. The primary goal is to create a scalable and efficient server that facilitates communication between multiple clients in real-time.

### Features

- **Real-time Communication**: Utilizing WebSocket for bidirectional communication between the server and clients, enabling instant message updates.
- **Scalability**: Designed to handle a large number of concurrent connections efficiently, ensuring smooth communication even with a high volume of users.
- **User Authentication**: Implement a secure authentication mechanism to validate users and manage access to chat features.
- **Message Persistence**: Optionally, store chat messages in a database to provide message history and continuity across user sessions.
- **Error Handling**: Implement robust error handling to gracefully manage unexpected scenarios, ensuring the stability of the chat server.
- **Logging**: Utilize logging to keep track of important events and errors for monitoring and debugging purposes.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/joaovds/gotterChat-backend
   ```

2. Navigate to the project directory:

   ```bash
   cd gotterChat-backend
   ```

3. Copy enviroments variables:

   ```bash
   cp .env.example .env
   ```

4. Install dependencies:

   ```bash
   go mod tidy
   ```

5. Run the server:

   ```bash
   go run cmd/chat/main.go
   ```

## Configuration

- Modify the `.env` file to customize the server settings, such as port, database connection details, and authentication configurations.

## Usage

1. Connect to the WebSocket server using the appropriate client library.
2. Implement user authentication based on the server's authentication mechanism.
3. Establish a WebSocket connection and start sending and receiving messages in real-time.

## Contributing

Contributions are welcome! Feel free to open issues for bug reports or feature requests. If you'd like to contribute, please fork the repository and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
