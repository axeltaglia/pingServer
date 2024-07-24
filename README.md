# PingServer

PingServer is a simple HTTP API server written in Go. It provides a single endpoint `/ping` which responds with a JSON message.

## Features

- `/ping` endpoint that returns a `Pong` message in JSON format.
- Structured logging using `slog` package.
- Graceful error handling and JSON response for errors.

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
   git clone https://github.com/yourusername/pingserver.git
   cd pingserver

2. Build the project:
   go build -o pingserver

## Usage

1. Run the server:
   ./pingserver

2. Access the `/ping` endpoint:
   Open your browser or use `curl` to access the endpoint:
   curl http://localhost:8081/ping

   You should get a response similar to:
   {
   "id": 0,
   "msg": "Pong"
   }
