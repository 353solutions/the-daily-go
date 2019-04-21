package main

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"nlp"
	"nlp/pb"
)

// NLPServer is NLP gRPC server
type NLPServer struct{}

// Tokenize is gRPC tokenization
func (s NLPServer) Tokenize(ctx context.Context, req *pb.TokenizeRequest) (*pb.TokenizeResponse, error) {
	resp := pb.TokenizeResponse{
		Tokens: nlp.Tokenize(req.Text),
	}

	return &resp, nil
}

func gRPCListenAndServe(addr string) error {
	srv := grpc.NewServer()
	pb.RegisterNLPServer(srv, NLPServer{})
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrapf(err, "can't listen on %s", addr)
	}

	return srv.Serve(lis)
}
