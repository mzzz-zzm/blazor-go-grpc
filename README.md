# blazor-go-grpc

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
