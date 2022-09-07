package auth

import (
	"github.com/aulyarahman/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/aulyarahman/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServieClient(c),
	}


	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}


func (svc *ServiceClient) Register(ctx *gin.Context) {
	 routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	 routes.Login(ctx, svc.Client)
}