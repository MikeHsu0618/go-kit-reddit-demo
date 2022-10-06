// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "auth/pkg/endpoint"
	pb "auth/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	generateToken grpc.Handler
	validateToken grpc.Handler
	refreshToken  grpc.Handler
	pb.UnimplementedAuthServer
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.AuthServer {
	return &grpcServer{
		generateToken: makeGenerateTokenHandler(endpoints, options["GenerateToken"]),
		refreshToken:  makeRefreshTokenHandler(endpoints, options["RefreshToken"]),
		validateToken: makeValidateTokenHandler(endpoints, options["ValidateToken"]),
	}
}