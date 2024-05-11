package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"petchain/config"
	"petchain/protos"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
)

type registrarServer struct {
	protos.UnimplementedRegistrarServer
}

func (server *registrarServer) RegisterDid(ctx context.Context, req *protos.RegistrarRequest) (*protos.RegistrarResponse, error) {
	log.Printf("Register DID: %s\n", req.Did)
	log.Printf("Register DID Document: %s\n", req.DidDocument)

	db, err := leveldb.OpenFile("did_db/dids", nil)
	if err != nil {
		panic(err)
	}
	defer func(db *leveldb.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err = db.Put([]byte(req.Did), []byte(req.DidDocument), nil)

	return &protos.RegistrarResponse{Result: "OK"}, nil
}

func main() {
	fmt.Println("### Start Registrar ###")
	lis, err := net.Listen("tcp", config.SystemConfig.RegistrarAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := registrarServer{}
	s := grpc.NewServer()
	protos.RegisterRegistrarServer(s, &server)

	log.Printf("Registrar Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
