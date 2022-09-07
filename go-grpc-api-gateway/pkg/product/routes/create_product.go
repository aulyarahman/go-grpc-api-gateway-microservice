package routes

import (
	"context"
	"net/http"

	"github.com/aulyarahman/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Stock int64  `json:"stock"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name: body.Name,
		Price: body.Price,
		Stock: body.Stock,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}