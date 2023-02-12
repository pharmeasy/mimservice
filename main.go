package main

import (
	"fmt"
	"github.com/pharmeasy/ni-go-boiler-plate/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	log.Printf("Initializing Config")
	config.Initialize()
	// create new gRPC server
	server := grpc.NewServer()
	// register your proto implementations
	//proto.RegisterAuthServiceServer(server, &auth.ServiceImpl{})
	// start listening on GrpcPort loaded in config
	reflection.Register(server)
	log.Println("GRPC is running on PORT:", config.Config.Env.GrpcPort)
	if l, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Config.Env.GrpcPort)); err != nil {
		log.Fatal("error in listening ", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
