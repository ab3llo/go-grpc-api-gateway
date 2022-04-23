package routes

import (
	"context"
	"net/http"

	"github.com/ab3llo/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := ctx.Params.Get("id")

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
