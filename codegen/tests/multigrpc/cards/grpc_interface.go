// Code generated by sysl DO NOT EDIT.
package cards

import (
	"context"

	pb "github.com/anz-bank/sysl-go/codegen/tests/multigrpc/cardspb"
)

// GetCards Client
type GetCardsClient struct {
}

// GrpcServiceInterface for Cards
type GrpcServiceInterface struct {
	GetCards func(ctx context.Context, req *pb.GetCardsRequest, client GetCardsClient) (*pb.GetCardsResponse, error)
}
