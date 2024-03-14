package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "") //local ip and port
	if err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}

	//初始化 gRpc server
	server := grpc.NewServer()
	// grpc 函数
	// func(server, &gRpcServer.)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}
}
