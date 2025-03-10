# blazor-go-grpc

```plaintext
blazor-go-grpc/
├── backend/                 # Go gRPC server
│   ├── main.go              # Main Go server code
│   ├── Dockerfile           # Go server Docker configuration
│   └── go.mod               # Go dependencies
├── frontend/                # Blazor WebAssembly client
│   ├── BlazorGrpcClient/    # Blazor project files
│   └── Dockerfile           # Blazor Docker configuration
├── proto/                   # Shared protocol buffer definitions
│   └── greet.proto          # gRPC service definition
├── docker-compose.yml       # Docker configuration for the entire app
└── README.md                # Setup instructions



# Blazor WebAssembly with Go gRPC Demo

This is a complete working example of a Blazor WebAssembly application communicating with a Go backend via gRPC-Web.

## Project Structure

- `backend/`: Go gRPC server
- `frontend/`: Blazor WebAssembly client
- `proto/`: Shared protocol buffer definitions

## Prerequisites

- Docker and Docker Compose
- .NET 8 SDK (for local development)
- Go 1.21+ (for local development)
- Protocol Buffers compiler

## Running with Docker

1. Clone this repository
2. Start the application with Docker Compose:

```bash
docker-compose up
```

3. Open your browser and navigate to http://localhost:5001

## Local Development Setup

### Backend (Go)

1. Navigate to the backend directory:

```bash
cd backend
```

2. Generate the Protocol Buffer files:

```bash
protoc -I../proto --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ../proto/greet.proto
```

3. Run the Go server:

```bash
go run main.go
```

### Frontend (Blazor WebAssembly)

1. Navigate to the frontend directory:

```bash
cd frontend/BlazorGrpcClient
```

2. The Protobuf files are generated automatically when building the project.

3. Run the Blazor app:

```bash
dotnet run
```

## How it Works

1. The gRPC service is defined in the `proto/greet.proto` file.
2. The Go backend implements this service and exposes it via gRPC-Web.
3. The Blazor WebAssembly client connects to the Go backend using the gRPC-Web protocol.

## Features

- Basic request-response communication
- Server-streaming demonstration
- Cross-platform compatibility

## Troubleshooting

- If you encounter CORS issues, check the CORS configuration in the Go backend.
- For networking issues in Docker, check that the service names match in the docker-compose.yml and Program.cs.
- Ensure the Protocol Buffer files are generated correctly for both Go and C#.