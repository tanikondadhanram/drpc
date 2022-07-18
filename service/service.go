package service

import (
	"context"
	"fmt"
	"net"

	DB "drpc/db"
	"drpc/pb"

	"google.golang.org/protobuf/types/known/structpb"
	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"
)

type Server struct {
	pb.DRPCStorageUnimplementedServer
	// struct fields
}

func (s *Server) GetData(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{ResponseData: &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"name": {Kind: &structpb.Value_StringValue{StringValue: "rk"}},
			"age":  {Kind: &structpb.Value_NumberValue{NumberValue: 22}},
		},
	}}, nil
}

func StartServer() {
	err := createServer(context.Background())
	if err != nil {
		panic(err)
	}
}

func createServer(ctx context.Context) error {
	DB.LoadDB()
	// create an RPC server
	server := &Server{}

	// create a drpc RPC mux
	m := drpcmux.New()

	// register the proto-specific methods on the mux
	err := pb.DRPCRegisterStorage(m, server)
	if err != nil {
		return err
	}

	// create a drpc server
	s := drpcserver.New(m)

	// listen on a tcp socket
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	fmt.Println("Starting service on 0.0.0.0:9000")
	// run the server
	// N.B.: if you want TLS, you need to wrap the net.Listener with
	// TLS before passing to Serve here.
	return s.Serve(ctx, lis)
}
