// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// protoc-gen-go-micro version: v3.5.3
// source: grpc.proto

package micro_old_server

import (
	context "context"
	api "go.unistack.org/micro/v3/api"
	client "go.unistack.org/micro/v3/client"
	server "go.unistack.org/micro/v3/server"
)

type microOldServerClient struct {
	c    client.Client
	name string
}

func NewMicroOldServerClient(name string, c client.Client) MicroOldServerClient {
	return &microOldServerClient{c: c, name: name}
}

func (c *microOldServerClient) Test(ctx context.Context, req *TestRequest, opts ...client.CallOption) (*TestResponse, error) {
	rsp := &TestResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "MicroOldServer.Test", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type microOldServerServer struct {
	MicroOldServerServer
}

func (h *microOldServerServer) Test(ctx context.Context, req *TestRequest, rsp *TestResponse) error {
	return h.MicroOldServerServer.Test(ctx, req, rsp)
}

func RegisterMicroOldServerServer(s server.Server, sh MicroOldServerServer, opts ...server.HandlerOption) error {
	type microOldServer interface {
		Test(ctx context.Context, req *TestRequest, rsp *TestResponse) error
	}
	type MicroOldServer struct {
		microOldServer
	}
	h := &microOldServerServer{sh}
	var nopts []server.HandlerOption
	for _, endpoint := range MicroOldServerEndpoints {
		nopts = append(nopts, api.WithEndpoint(&endpoint))
	}
	return s.Handle(s.NewHandler(&MicroOldServer{h}, append(nopts, opts...)...))
}
