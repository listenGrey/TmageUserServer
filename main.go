package main

import (
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"google.golang.org/grpc"
	"log"
	"net"

	local "TmageUsersServer/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8964") //local ip and port
	if err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}

	//初始化 gRpc server
	server := grpc.NewServer()

	userInfo.RegisterCheckExistenceServer(server, &local.ExistenceServer{})
	userInfo.RegisterRegisterInfoServer(server, &local.RegisterServer{})
	userInfo.RegisterLoginCheckServer(server, &local.LoginServer{})

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to connect, %s", err)
	}
}
