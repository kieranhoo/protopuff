package gateway

import (
	"context"
	"fmt"
	"net"

	"protopuff/internal/gen/v1/greeter"
	"protopuff/internal/module/service"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	gate *gin.Engine
	http string
	grpc string
}

func New(httpUri string, grpcUri string) *Gateway {
	gin.SetMode(gin.ReleaseMode)
	log.Info("Server registered with the following settings:")
	log.Info("- HTTP", "[TCP]", httpUri)
	log.Info("- gRPC", "[TCP]", grpcUri)
	fmt.Println()
	return &Gateway{
		gate: gin.Default(),
		http: httpUri,
		grpc: grpcUri,
	}
}

func (g *Gateway) prepareHttpServer() *Gateway {
	GinMiddleware(g.gate)
	return g
}

func (g *Gateway) run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := greeter.RegisterGreeterHandlerFromEndpoint(ctx, mux, g.grpc, opts)
	if err != nil {
		return err
	}

	g.gate.Any("/*any", gin.WrapF(mux.ServeHTTP))
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return g.gate.Run(g.http)
	// return http.ListenAndServe(":8081", mux)
}

func (g *Gateway) startGrpcServer() *Gateway {
	listener, err := net.Listen("tcp", g.grpc)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServer(grpcServer, service.NewGreeter())

	fmt.Println()
	go grpcServer.Serve(listener)
	return g
}

func (g *Gateway) Serve() error {
	return g.prepareHttpServer().
		startGrpcServer().
		run(context.Background())
}
