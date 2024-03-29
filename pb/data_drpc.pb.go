// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.32
// source: pb/data.proto

package pb

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_pb_data_proto struct{}

func (drpcEncoding_File_pb_data_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_pb_data_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_pb_data_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_pb_data_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_pb_data_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCStorageClient interface {
	DRPCConn() drpc.Conn

	GetData(ctx context.Context, in *Request) (*Response, error)
}

type drpcStorageClient struct {
	cc drpc.Conn
}

func NewDRPCStorageClient(cc drpc.Conn) DRPCStorageClient {
	return &drpcStorageClient{cc}
}

func (c *drpcStorageClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcStorageClient) GetData(ctx context.Context, in *Request) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/data.Storage/GetData", drpcEncoding_File_pb_data_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCStorageServer interface {
	GetData(context.Context, *Request) (*Response, error)
}

type DRPCStorageUnimplementedServer struct{}

func (s *DRPCStorageUnimplementedServer) GetData(context.Context, *Request) (*Response, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCStorageDescription struct{}

func (DRPCStorageDescription) NumMethods() int { return 1 }

func (DRPCStorageDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/data.Storage/GetData", drpcEncoding_File_pb_data_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCStorageServer).
					GetData(
						ctx,
						in1.(*Request),
					)
			}, DRPCStorageServer.GetData, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterStorage(mux drpc.Mux, impl DRPCStorageServer) error {
	return mux.Register(impl, DRPCStorageDescription{})
}

type DRPCStorage_GetDataStream interface {
	drpc.Stream
	SendAndClose(*Response) error
}

type drpcStorage_GetDataStream struct {
	drpc.Stream
}

func (x *drpcStorage_GetDataStream) SendAndClose(m *Response) error {
	if err := x.MsgSend(m, drpcEncoding_File_pb_data_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
