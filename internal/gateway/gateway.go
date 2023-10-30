package gateway

import (
	"context"
	"flag"

	"protopuff/internal/gen/v1/greeter"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	gate               *gin.Engine
	grpcServerEndpoint *string
	http               string
	grpc               string
}

func New(httpUri string, grpcUri string) *Gateway {
	return &Gateway{
		gate: gin.Default(),
		// command-line options:
		// gRPC server endpoint
		grpcServerEndpoint: flag.String("grpc-server-endpoint", grpcUri, "gRPC server endpoint"),
		http:               httpUri,
		grpc:               grpcUri,
	}
}

func (g *Gateway) run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := greeter.RegisterGreeterHandlerFromEndpoint(ctx, mux, *g.grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	log.Info("grpc server registered with the following settings:")
	log.Info("- listen", "[TCP]", g.grpc)
	g.gate.Any("/*any", gin.WrapF(mux.ServeHTTP))
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return g.gate.Run(g.http)
	// return http.ListenAndServe(":8081", mux)
}

func (g *Gateway) Serve() error {
	return g.run()
}
