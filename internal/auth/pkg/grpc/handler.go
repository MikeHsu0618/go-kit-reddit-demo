package grpc

import (
	endpoint "auth/pkg/endpoint"
	pb "auth/pkg/grpc/pb"
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// makeGenerateTokenHandler creates the handler logic
func makeGenerateTokenHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GenerateTokenEndpoint, decodeGenerateTokenRequest, encodeGenerateTokenResponse, options...)
}

// decodeGenerateTokenResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GenerateToken request.
// TODO implement the decoder
func decodeGenerateTokenRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Decoder is not impelemented")
}

// encodeGenerateTokenResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGenerateTokenResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Encoder is not impelemented")
}
func (g *grpcServer) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenReply, error) {
	_, rep, err := g.generateToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenerateTokenReply), nil
}

// makeValidateTokenHandler creates the handler logic
func makeValidateTokenHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ValidateTokenEndpoint, decodeValidateTokenRequest, encodeValidateTokenResponse, options...)
}

// decodeValidateTokenResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ValidateToken request.
// TODO implement the decoder
func decodeValidateTokenRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Decoder is not impelemented")
}

// encodeValidateTokenResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeValidateTokenResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Encoder is not impelemented")
}
func (g *grpcServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenReply, error) {
	_, rep, err := g.validateToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ValidateTokenReply), nil
}

// makeRefreshTokenHandler creates the handler logic
func makeRefreshTokenHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RefreshTokenEndpoint, decodeRefreshTokenRequest, encodeRefreshTokenResponse, options...)
}

// decodeRefreshTokenResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain RefreshToken request.
// TODO implement the decoder
func decodeRefreshTokenRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Decoder is not impelemented")
}

// encodeRefreshTokenResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRefreshTokenResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Auth' Encoder is not impelemented")
}
func (g *grpcServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	_, rep, err := g.refreshToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RefreshTokenReply), nil
}
