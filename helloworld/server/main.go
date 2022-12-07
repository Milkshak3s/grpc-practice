package main

import (
	pb "github.com/Milkshak3s/helloworld/helloworld"
)

type server struct {
	pb.UnimplementedGreeterServer
}
