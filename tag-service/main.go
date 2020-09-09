package main

import (
	pb "github.com/noChaos1012/tour/tag-service/proto"
	"github.com/noChaos1012/tour/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port = "8999"

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())

	reflection.Register(s) //注册反射服务，以备grpcurls使用

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("net.Listen err:%v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("server.Serve err:%v", err)
	}
}
