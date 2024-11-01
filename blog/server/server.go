package main

import pb "grpc-crud/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
