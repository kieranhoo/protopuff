package main

import (
	"context"
	"log"
	"protopuff/internal/config"
	"protopuff/pkg/gen/v1/greeter"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	address string
	conn    *grpc.ClientConn
}

func New() *Client {
	conn, err := grpc.Dial(config.RpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &Client{
		address: config.RpcHost,
		conn:    conn,
	}
}

func (c *Client) SayHelloService() error {
	defer c.CloseConnection()
	client := greeter.NewGreeterClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &greeter.MessageRequest{
		Message:   "Hello from service provider",
		Timestamp: time.Now().UTC().Unix(),
	})
	if err != nil {
		return err
	}
	log.Printf("Greeting: [%s] at timestamp - [%v]", r.GetMessage(), r.GetTimestamp())
	return nil
}

func (c *Client) CloseConnection() {
	err := c.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client := New()
	client.SayHelloService()
}
