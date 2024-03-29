// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Reservation service

type ReservationService interface {
	Reserve(ctx context.Context, in *ReservationRequest, opts ...client.CallOption) (*ReservationResponse, error)
	AcceptReservation(ctx context.Context, in *AcceptReservationRequest, opts ...client.CallOption) (*AcceptReservationResponse, error)
	Cancel(ctx context.Context, in *CancelReservationRequest, opts ...client.CallOption) (*CancelReservationResponse, error)
	CancelForPresentations(ctx context.Context, in *CancelForPresentationsRequest, opts ...client.CallOption) (*CancelForPresentationsResponse, error)
	CancelForUsers(ctx context.Context, in *CancelForUsersRequest, opts ...client.CallOption) (*CancelForUsersResponse, error)
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...client.CallOption) (*ReadAllResponse, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...client.CallOption) (*ClearResponse, error)
}

type reservationService struct {
	c    client.Client
	name string
}

func NewReservationService(name string, c client.Client) ReservationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &reservationService{
		c:    c,
		name: name,
	}
}

func (c *reservationService) Reserve(ctx context.Context, in *ReservationRequest, opts ...client.CallOption) (*ReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.Reserve", in)
	out := new(ReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) AcceptReservation(ctx context.Context, in *AcceptReservationRequest, opts ...client.CallOption) (*AcceptReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.AcceptReservation", in)
	out := new(AcceptReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) Cancel(ctx context.Context, in *CancelReservationRequest, opts ...client.CallOption) (*CancelReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.Cancel", in)
	out := new(CancelReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) CancelForPresentations(ctx context.Context, in *CancelForPresentationsRequest, opts ...client.CallOption) (*CancelForPresentationsResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.CancelForPresentations", in)
	out := new(CancelForPresentationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) CancelForUsers(ctx context.Context, in *CancelForUsersRequest, opts ...client.CallOption) (*CancelForUsersResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.CancelForUsers", in)
	out := new(CancelForUsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...client.CallOption) (*ReadAllResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.ReadAll", in)
	out := new(ReadAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) Clear(ctx context.Context, in *ClearRequest, opts ...client.CallOption) (*ClearResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.Clear", in)
	out := new(ClearResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Reservation service

type ReservationHandler interface {
	Reserve(context.Context, *ReservationRequest, *ReservationResponse) error
	AcceptReservation(context.Context, *AcceptReservationRequest, *AcceptReservationResponse) error
	Cancel(context.Context, *CancelReservationRequest, *CancelReservationResponse) error
	CancelForPresentations(context.Context, *CancelForPresentationsRequest, *CancelForPresentationsResponse) error
	CancelForUsers(context.Context, *CancelForUsersRequest, *CancelForUsersResponse) error
	ReadAll(context.Context, *ReadAllRequest, *ReadAllResponse) error
	Clear(context.Context, *ClearRequest, *ClearResponse) error
}

func RegisterReservationHandler(s server.Server, hdlr ReservationHandler, opts ...server.HandlerOption) error {
	type reservation interface {
		Reserve(ctx context.Context, in *ReservationRequest, out *ReservationResponse) error
		AcceptReservation(ctx context.Context, in *AcceptReservationRequest, out *AcceptReservationResponse) error
		Cancel(ctx context.Context, in *CancelReservationRequest, out *CancelReservationResponse) error
		CancelForPresentations(ctx context.Context, in *CancelForPresentationsRequest, out *CancelForPresentationsResponse) error
		CancelForUsers(ctx context.Context, in *CancelForUsersRequest, out *CancelForUsersResponse) error
		ReadAll(ctx context.Context, in *ReadAllRequest, out *ReadAllResponse) error
		Clear(ctx context.Context, in *ClearRequest, out *ClearResponse) error
	}
	type Reservation struct {
		reservation
	}
	h := &reservationHandler{hdlr}
	return s.Handle(s.NewHandler(&Reservation{h}, opts...))
}

type reservationHandler struct {
	ReservationHandler
}

func (h *reservationHandler) Reserve(ctx context.Context, in *ReservationRequest, out *ReservationResponse) error {
	return h.ReservationHandler.Reserve(ctx, in, out)
}

func (h *reservationHandler) AcceptReservation(ctx context.Context, in *AcceptReservationRequest, out *AcceptReservationResponse) error {
	return h.ReservationHandler.AcceptReservation(ctx, in, out)
}

func (h *reservationHandler) Cancel(ctx context.Context, in *CancelReservationRequest, out *CancelReservationResponse) error {
	return h.ReservationHandler.Cancel(ctx, in, out)
}

func (h *reservationHandler) CancelForPresentations(ctx context.Context, in *CancelForPresentationsRequest, out *CancelForPresentationsResponse) error {
	return h.ReservationHandler.CancelForPresentations(ctx, in, out)
}

func (h *reservationHandler) CancelForUsers(ctx context.Context, in *CancelForUsersRequest, out *CancelForUsersResponse) error {
	return h.ReservationHandler.CancelForUsers(ctx, in, out)
}

func (h *reservationHandler) ReadAll(ctx context.Context, in *ReadAllRequest, out *ReadAllResponse) error {
	return h.ReservationHandler.ReadAll(ctx, in, out)
}

func (h *reservationHandler) Clear(ctx context.Context, in *ClearRequest, out *ClearResponse) error {
	return h.ReservationHandler.Clear(ctx, in, out)
}
