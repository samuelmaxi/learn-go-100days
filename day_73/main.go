package main

import (
	pb "github.com/aprendagolag/classificados/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.NewClient("localhost:9000", creds)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewC
}
