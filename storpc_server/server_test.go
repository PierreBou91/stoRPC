package main

import (
	"context"
	"log"
	"testing"

	pb "github.com/PierreBou91/stoRPC/storpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const TEST_HOST string = ":8080"

func storpcClientForTest(host string) (pb.StorpcClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return pb.NewStorpcClient(conn), err
}
func TestServices(t *testing.T) {
	// test setup:
	// 1. launch server
	// 2. create client
	// 3. define test variables
	// 4. run test cases

	// 1. launch server
	go launchServer(TEST_HOST) // possible problem where server is not ready yet

	ctx := context.Background()

	// 2. create client
	client, err := storpcClientForTest(TEST_HOST)
	if err != nil {
		t.Errorf("Error while creating client: %v", err)
	}

	// 3. define test variables
	pair := &pb.Pair{Key: "key", Value: "value"}

	log.Printf("Running tests against server at %s", TEST_HOST)
	t.Run("User can store a key value pair on server", func(t *testing.T) {
		res, err := client.PutValue(ctx, pair)
		if err != nil || !res.Ok {
			t.Errorf("Error while calling PutValue: %v", err)
		}
	})

	t.Run("User can retrieve a value from the server", func(t *testing.T) {
		res, err := client.GetValue(ctx, &pb.Key{Key: pair.Key})
		if err != nil || res.Value != pair.Value {
			t.Errorf("Error while calling GetValue: %v", err)
		}
	})

	t.Run("User can delete a key value pair from the server", func(t *testing.T) {
		res, err := client.DeleteValue(ctx, &pb.Key{Key: pair.Key})
		if err != nil || !res.Ok {
			t.Errorf("Error while calling DeleteValue: %v", err)
		}
	})

}
