package auth

import (
	"fmt"

	"github.com/ab3llo/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/ab3llo/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect", err)
	}
	return pb.NewAuthServiceClient(cc)
}
