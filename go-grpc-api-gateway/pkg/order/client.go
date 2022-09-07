package order

import (
	"fmt"

	"github.com/aulyarahman/go-grpc-api-gateway/pkg/config"
	"github.com/aulyarahman/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}


func InitServieClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect to service :", err)
	}

	return pb.NewOrderServiceClient(cc)
}