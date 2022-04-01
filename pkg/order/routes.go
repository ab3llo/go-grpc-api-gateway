package order

import (
	"github.com/ab3llo/go-grpc-api-gateway/pkg/auth"
	"github.com/ab3llo/go-grpc-api-gateway/pkg/config"
	"github.com/ab3llo/go-grpc-api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/api/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
