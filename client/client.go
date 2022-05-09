package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "gRPC/assignment/test_proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSimpleMathClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//var n1, n2 int32
	//fmt.Println("enter the first number:")
	//n1,_ = fmt.Scanf()
	//fmt.Println("enter the second number:")
	//fmt.Scanf()
	if os.Args[1] == "add" {
		r, err := c.Add(ctx, &pb.Number{Num1: 55, Num2: 11})
		if err != nil {
			log.Fatalf("could not complete : %v", err)
		}
		log.Printf(`OUTPUT: %d`, r.Out)
	} else if os.Args[1] == "sub" {
		r, err := c.Sub(ctx, &pb.Number{Num1: 55, Num2: 11})
		if err != nil {
			log.Fatalf("could not complete : %v", err)
		}
		log.Printf(`OUTPUT: %d`, r.Out)
	} else if os.Args[1] == "mul" {
		r, err := c.Mul(ctx, &pb.Number{Num1: 55, Num2: 11})
		if err != nil {
			log.Fatalf("could not complete : %v", err)
		}
		log.Printf(`OUTPUT: %d`, r.Out)
	} else if os.Args[1] == "div" {
		r, err := c.Div(ctx, &pb.Number{Num1: 55, Num2: 11})
		if err != nil {
			log.Fatalf("could not complete : %v", err)
		}
		log.Printf(`OUTPUT: %d`, r.Out)
	}

}
