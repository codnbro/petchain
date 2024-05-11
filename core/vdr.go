package core

import (
	"context"
	"errors"
	"fmt"
	"log"
	"petchain/config"
	"petchain/protos"
	"time"

	"google.golang.org/grpc"
)

func RegisterDid(did string, didDocument string) error {
	conn, err := grpc.Dial(config.SystemConfig.RegistrarAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Registrar not connect: %v\n", err)
		return errors.New(fmt.Sprintf("Registrar not connect: %v", err))
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := protos.NewRegistrarClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.RegisterDid(ctx, &protos.RegistrarRequest{Did: did, DidDocument: didDocument})
	if err != nil {
		log.Println("Failed to register DID.")
		return errors.New("failed to register DID")
	}

	fmt.Printf("Registrar Response: %s\n", res)

	return nil
}

func ResolveDid(did string) (string, error) {
	conn, err := grpc.Dial(config.SystemConfig.ResolverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Resolver not connect: %v\n", err)
		return "", errors.New(fmt.Sprintf("Resolver not connect: %v", err))
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := protos.NewResolverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.ResolveDid(ctx, &protos.ResolverRequest{Did: did})
	if err != nil {
		log.Fatalf("Failed to resolve DID.")
	}

	fmt.Printf("Result: %s\n", res)

	return res.DidDocument, nil
}
