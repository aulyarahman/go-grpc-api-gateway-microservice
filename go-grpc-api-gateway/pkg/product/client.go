package product

import (
	"fmt"

	"github.com/aulyarahman/go-grpc-api-gateway/pkg/config"
	"github.com/aulyarahman/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServieClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl,  grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect to service :", err)
	}

	return pb.NewProductServiceClient(cc)
}