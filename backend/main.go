package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/mzzz-zzm/blazor-go-grpc/backend/pb"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements the SayHello RPC method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request from: %s", req.GetName())
	return &pb.HelloReply{
		Message:   fmt.Sprintf("Hello, %s!", req.GetName()),
		Timestamp: time.Now().Format(time.RFC3339),
	}, nil
}

// SayHellos implements the streaming SayHellos RPC method
func (s *server) SayHellos(req *pb.HelloRequest, stream pb.Greeter_SayHellosServer) error {
	log.Printf("Received streaming request from: %s", req.GetName())

	for i := 0; i < 5; i++ {
		// Send multiple greetings
		err := stream.Send(&pb.HelloReply{
			Message:   fmt.Sprintf("Hello %d to %s!", i+1, req.GetName()),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Second) // Delay between messages
	}
	return nil
}

func main() {
	// Create a standard gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	// Create a gRPC-Web wrapper for the gRPC server
	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true // Allow all origins for demo purposes
		}),
		grpcweb.WithWebsockets(true),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
	)

	// Create a CORS handler
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins for demo
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow all headers
		AllowCredentials: true,
		ExposedHeaders:   []string{"grpc-status", "grpc-message"}, // Expose gRPC headers
	})

	// Create an HTTP handler that wraps the gRPC-Web handler with CORS support
	httpHandler := corsHandler.Handler(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("Received %s request for: %s", req.Method, req.URL.Path)

		// Handle gRPC-Web requests or CORS preflight requests
		if wrappedGrpc.IsGrpcWebRequest(req) || wrappedGrpc.IsAcceptableGrpcCorsRequest(req) {
			log.Println("Handling as gRPC-Web request")
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}

		// Handle regular HTTP requests (helpful for checking server status)
		if req.URL.Path == "/health" {
			resp.WriteHeader(http.StatusOK)
			resp.Write([]byte("Server is healthy"))
			return
		}

		// Default root handler
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("gRPC-Web server is running. Use a gRPC-Web client to connect."))
	}))

	// Start the server
	port := 8080
	log.Printf("Starting gRPC-Web server on port %d...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), httpHandler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
