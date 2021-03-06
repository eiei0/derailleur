// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/derailleur.proto

package derailleur

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Derailleur service

func NewDerailleurEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Derailleur service

type DerailleurService interface {
	CreateDerailleur(ctx context.Context, in *CreateDerailleurRequest, opts ...client.CallOption) (*CreateDerailleurResponse, error)
}

type derailleurService struct {
	c    client.Client
	name string
}

func NewDerailleurService(name string, c client.Client) DerailleurService {
	return &derailleurService{
		c:    c,
		name: name,
	}
}

func (c *derailleurService) CreateDerailleur(ctx context.Context, in *CreateDerailleurRequest, opts ...client.CallOption) (*CreateDerailleurResponse, error) {
	req := c.c.NewRequest(c.name, "Derailleur.CreateDerailleur", in)
	out := new(CreateDerailleurResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Derailleur service

type DerailleurHandler interface {
	CreateDerailleur(context.Context, *CreateDerailleurRequest, *CreateDerailleurResponse) error
}

func RegisterDerailleurHandler(s server.Server, hdlr DerailleurHandler, opts ...server.HandlerOption) error {
	type derailleur interface {
		CreateDerailleur(ctx context.Context, in *CreateDerailleurRequest, out *CreateDerailleurResponse) error
	}
	type Derailleur struct {
		derailleur
	}
	h := &derailleurHandler{hdlr}
	return s.Handle(s.NewHandler(&Derailleur{h}, opts...))
}

type derailleurHandler struct {
	DerailleurHandler
}

func (h *derailleurHandler) CreateDerailleur(ctx context.Context, in *CreateDerailleurRequest, out *CreateDerailleurResponse) error {
	return h.DerailleurHandler.CreateDerailleur(ctx, in, out)
}
