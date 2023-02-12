package main

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/PierreBou91/stoRPC/storpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
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
	ctx := context.Background()

	// 1. launch server
	go launchServer(TEST_HOST)

	waitForTestServerToBeReady(ctx, t)

	// 2. create client
	client, err := storpcClientForTest(TEST_HOST)
	if err != nil {
		t.Errorf("Error while creating client: %v", err)
	}

	// 3. define test variables
	pair := &pb.Pair{Key: "key", Value: "value"}

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

// waitForServerToBeReady is a UGLY and should be corrected to implement
// a proper healthcheck that times out with the context
func waitForTestServerToBeReady(ctx context.Context, t *testing.T) {
	healthConn, err := grpc.Dial(TEST_HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Errorf("Error while creating health client: %v", err)
	}

	healthClient := healthpb.NewHealthClient(healthConn)

	// UGLY
	for {
		resp, err := healthClient.Check(ctx, &healthpb.HealthCheckRequest{})
		if err != nil {
			t.Errorf("Error while calling health check: %v", err)
		}

		if resp.Status == healthpb.HealthCheckResponse_SERVING {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}
