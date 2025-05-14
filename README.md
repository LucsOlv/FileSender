# FileSender Microservice

A microservice for uploading files and sending them to a RabbitMQ queue for further processing.

## Overview

FileSender is a Go-based microservice that provides a REST API for file uploads. When files are uploaded, they are published to a RabbitMQ queue, allowing other services to process them asynchronously. This architecture enables scalable file processing in distributed systems.

## Features

- RESTful API for file uploads
- Multiple file upload support
- Asynchronous file processing via RabbitMQ
- Swagger documentation
- Graceful server shutdown
- CORS support

## Prerequisites

- Go 1.16 or higher
- RabbitMQ server
- Docker (optional, for containerization)

## Installation

### Clone the repository

```bash
git clone https://github.com/LucsOlv/FileSender.git
cd filesender
```

### Install dependencies

```bash
go mod download
```

## Configuration

The application uses environment variables for configuration. Create a `.env` file in the root directory with the following variables:

```
SERVER_PORT=8080
RABBITMQ_URI=amqp://guest:guest@localhost:5672/
```

## Running the Application

### Start RabbitMQ

Make sure RabbitMQ is running. You can start it using Docker:

```bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

### Run the application

```bash
go run main.go
```

The server will start on the configured port (default: 8080).

## API Documentation

Swagger documentation is available at `/swagger/index.html` when the server is running.

### Endpoints

- `GET /ping` - Health check endpoint
- `POST /api/form` - Upload files endpoint

### Upload Files

To upload files, send a POST request to `/api/form` with the files in the `files` field of a multipart/form-data request.

Example using curl:

```bash
curl -X POST http://localhost:8080/api/form \
  -F "files=@/path/to/file1.txt" \
  -F "files=@/path/to/file2.jpg"
```

## Project Structure

```
├── config/         # Configuration loading
├── docs/           # Swagger documentation
├── handlers/       # HTTP request handlers
├── messaging/      # RabbitMQ integration
├── routes/         # API route definitions
├── main.go         # Application entry point
└── README.md       # This file
```

## RabbitMQ Message Format

Files are published to the `file_queue` queue with the following properties:

- Content-Type: application/octet-stream
- Body: Raw file bytes
- Headers:
  - filename: Original filename

## Development

### Generate Swagger Documentation

To update the Swagger documentation, run:

```bash
swag init
```

## License

[MIT](LICENSE)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
