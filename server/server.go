package main

import (
	"context"
	"log"
	"net"

	pb "gRPC/assignment/test_proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type SimpleMathServer struct {
	pb.UnimplementedSimpleMathServer
}

func (s *SimpleMathServer) Add(ctx context.Context, in *pb.Number) (*pb.Out, error) {
	log.Printf("Numbers received for add: %v", in)
	var out int32 = in.Num1 + in.Num2
	return &pb.Out{Out: out}, nil
}

func (s *SimpleMathServer) Mul(ctx context.Context, in *pb.Number) (*pb.Out, error) {
	log.Printf("Numbers received for mul: %v", in)
	var out int32 = in.Num1 * in.Num2
	return &pb.Out{Out: out}, nil
}

func (s *SimpleMathServer) Div(ctx context.Context, in *pb.Number) (*pb.Out, error) {
	log.Printf("Numbers received for div : %v", in)
	var out int32
	if in.Num1 > in.Num2 {
		out = in.Num1 / in.Num2
	} else {
		out = in.Num2 / in.Num1
	}
	return &pb.Out{Out: out}, nil
}

func (s *SimpleMathServer) Sub(ctx context.Context, in *pb.Number) (*pb.Out, error) {
	log.Printf("Numbers received for sub : %v", in)
	var out int32
	if in.Num1 > in.Num2 {
		out = in.Num1 - in.Num2
	} else {
		out = in.Num2 - in.Num1
	}
	return &pb.Out{Out: out}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSimpleMathServer(s, &SimpleMathServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
