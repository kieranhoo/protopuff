package gateway

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"protopuff/internal/config"
	"protopuff/internal/mod/service"
	"protopuff/internal/proto/gen/v1/greeter"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var gwlogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    false,
	ReportTimestamp: true,
	TimeFormat:      time.Kitchen,
	Prefix:          "[PROTOPUFF]",
})

type Gateway struct {
	gate *gin.Engine
	http string
	grpc string
}

func New() *Gateway {
	gin.SetMode(gin.ReleaseMode)
	gwlogger.Info("Server registered with the following settings:")
	gwlogger.Info("- HTTP", "[TCP]", config.HttpHost)
	gwlogger.Info("- gRPC", "[TCP]", config.RpcHost)
	fmt.Println()
	return &Gateway{
		gate: gin.Default(),
		http: config.HttpHost,
		grpc: config.RpcHost,
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

	// Register service endpoint here
	if err := greeter.RegisterGreeterHandlerFromEndpoint(ctx, mux, g.grpc, opts); err != nil {
		return err
	}

	g.gate.Any("/*any", gin.WrapF(mux.ServeHTTP))
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return g.gate.Run(g.http)
	// return http.ListenAndServe(g.http, mux)
}

func (g *Gateway) startGrpcServer() *Gateway {
	listener, err := net.Listen("tcp", g.grpc)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(GrpcLogger))

	// Start gRPC server here
	greeter.RegisterGreeterServer(grpcServer, service.NewGreeter())

	go grpcServer.Serve(listener)
	return g
}

func (g *Gateway) ServeGateway() error {
	return g.prepareHttpServer().startGrpcServer().run(context.Background())
}

func (g *Gateway) ServeGrpcServer() {
	g.startGrpcServer()
}
