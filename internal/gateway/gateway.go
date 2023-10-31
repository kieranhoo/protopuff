package gateway

import (
	"context"
	"fmt"
	"net"

	"protopuff/internal/proto/gen/v1/greeter"

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

func (g *Gateway) startGrpcServer(rpc func(*grpc.Server)) *Gateway {
	listener, err := net.Listen("tcp", g.grpc)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	rpc(grpcServer)
	go grpcServer.Serve(listener)
	return g
}

// ServeGateway
// rpc Example:
//
//	func registerGrpcServer(s *grpc.Server) {
//		greeter.RegisterGreeterServer(s, service.NewGreeter())
//	}
func (g *Gateway) ServeGateway(rpc func(*grpc.Server)) error {
	return g.prepareHttpServer().startGrpcServer(rpc).run(context.Background())
}

func (g *Gateway) ServeGrpcServer(rpc func(*grpc.Server)) {
	g.startGrpcServer(rpc)
}
