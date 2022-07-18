package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"drpc/pb"

	"google.golang.org/protobuf/types/known/structpb"
	"storj.io/drpc/drpcconn"
)

func main() {
	err := RequestServer(context.Background())
	if err != nil {
		fmt.Println("error returned from RequestServer", err)
		panic(err)
	}
}

func RequestServer(ctx context.Context) error {
	// dial the drpc server
	rawconn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("error returned from net.Dial", err)
		return err
	}
	// N.B.: If you want TLS, you need to wrap the net.Conn with TLS before
	// making a DRPC conn.

	// convert the net.Conn to a drpc.Conn
	conn := drpcconn.New(rawconn)
	defer conn.Close()

	// make a drpc proto-specific client
	client := pb.NewDRPCStorageClient(conn)

	// set a deadline for the operation
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// run the RPC
	data, err := client.GetData(ctx, &pb.Request{
		RequestData: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"name": {Kind: &structpb.Value_StringValue{StringValue: "dhanush"}},
				"age":  {Kind: &structpb.Value_NumberValue{NumberValue: 20}},
			},
		},
	})
	if err != nil {
		fmt.Println("error returned from client.GetData", err)
		return err
	}

	// check the results
	_, err = fmt.Println(data)
	return err
}
