// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: plugins/catdog_debug/proto/debug.proto
package debug

import (
	context "context"
	fmt "fmt"
	math "math"

	"github.com/pubgo/catdog/catdog_data"

	client "github.com/asim/nitro/v3/client"
	server "github.com/asim/nitro/v3/server"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Client API for Debug service
type DebugService interface {
	Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error)
	Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error)
}

type debugService struct {
	c    client.Client
	name string
}

func NewDebugService(name string, c client.Client) DebugService {
	return &debugService{
		c:    c,
		name: name,
	}
}
func (c *DebugService) Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error) {

	req := c.c.NewRequest(c.name, "Debug.Log", in)
	out := new(LogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type Debug_LogService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type DebugLog struct {
	stream client.Stream
}

func (x *DebugLog) Close() error {
	return x.stream.Close()
}

func (x *DebugLog) Context() context.Context {
	return x.stream.Context()
}

func (x *DebugLog) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *DebugLog) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *DebugLog) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *DebugService) Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error) {

	req := c.c.NewRequest(c.name, "Debug.Health", in)
	out := new(HealthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type Debug_HealthService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type DebugHealth struct {
	stream client.Stream
}

func (x *DebugHealth) Close() error {
	return x.stream.Close()
}

func (x *DebugHealth) Context() context.Context {
	return x.stream.Context()
}

func (x *DebugHealth) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *DebugHealth) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *DebugHealth) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *DebugService) Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error) {

	req := c.c.NewRequest(c.name, "Debug.Stats", in)
	out := new(StatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type Debug_StatsService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type DebugStats struct {
	stream client.Stream
}

func (x *DebugStats) Close() error {
	return x.stream.Close()
}

func (x *DebugStats) Context() context.Context {
	return x.stream.Context()
}

func (x *DebugStats) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *DebugStats) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *DebugStats) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *DebugService) Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error) {

	req := c.c.NewRequest(c.name, "Debug.Trace", in)
	out := new(TraceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type Debug_TraceService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type DebugTrace struct {
	stream client.Stream
}

func (x *DebugTrace) Close() error {
	return x.stream.Close()
}

func (x *DebugTrace) Context() context.Context {
	return x.stream.Context()
}

func (x *DebugTrace) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *DebugTrace) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *DebugTrace) Send(m *Message) error {
	return x.stream.Send(m)
}

// Server API for Debug service
type DebugHandler interface {
	Log(context.Context, *LogRequest, *LogResponse) error
	Health(context.Context, *HealthRequest, *HealthResponse) error
	Stats(context.Context, *StatsRequest, *StatsResponse) error
	Trace(context.Context, *TraceRequest, *TraceResponse) error
}

func RegisterDebugHandler(s server.Server, hdlr DebugHandler, opts ...server.HandlerOption) error {
	type debug interface {
		Log(ctx context.Context, in *LogRequest, out *LogResponse) error
		Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error
		Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error
		Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error
	}

	type Debug struct {
		debug
	}
	h := &debugHandler{hdlr}
	opts = append(opts, server.EndpointMetadata("Log", map[string]string{"POST": "/log"}))
	opts = append(opts, server.EndpointMetadata("Health", map[string]string{"POST": "/health"}))
	opts = append(opts, server.EndpointMetadata("Stats", map[string]string{"POST": "/stats"}))
	opts = append(opts, server.EndpointMetadata("Trace", map[string]string{"POST": "/trace"}))
	return s.Handle(s.NewHandler(&Debug{h}, opts...))
}

func init() { catdog_data.Add("RegisterDebugHandler", RegisterDebugHandler) }

type debugHandler struct {
	DebugHandler
}

func (h *DebugHandler) Log(ctx context.Context, in *LogRequest, out *LogResponse) error {
	return h.DebugHandler.Log(ctx, in, out)
}

func (h *DebugHandler) Log(ctx context.Context, stream server.Stream) error {

	m := new(LogRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DebugHandler.Log(ctx, m, &debugLogStream{stream})

}

type Debug_LogStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type debugLogStream struct {
	stream server.Stream
}

func (x *debugLogStream) Close() error {
	return x.stream.Close()
}

func (x *debugLogStream) Context() context.Context {
	return x.stream.Context()
}

func (x *debugLogStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugLogStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *DebugHandler) Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error {
	return h.DebugHandler.Health(ctx, in, out)
}

func (h *DebugHandler) Health(ctx context.Context, stream server.Stream) error {

	m := new(HealthRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DebugHandler.Health(ctx, m, &debugHealthStream{stream})

}

type Debug_HealthStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type debugHealthStream struct {
	stream server.Stream
}

func (x *debugHealthStream) Close() error {
	return x.stream.Close()
}

func (x *debugHealthStream) Context() context.Context {
	return x.stream.Context()
}

func (x *debugHealthStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugHealthStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *DebugHandler) Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error {
	return h.DebugHandler.Stats(ctx, in, out)
}

func (h *DebugHandler) Stats(ctx context.Context, stream server.Stream) error {

	m := new(StatsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DebugHandler.Stats(ctx, m, &debugStatsStream{stream})

}

type Debug_StatsStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type debugStatsStream struct {
	stream server.Stream
}

func (x *debugStatsStream) Close() error {
	return x.stream.Close()
}

func (x *debugStatsStream) Context() context.Context {
	return x.stream.Context()
}

func (x *debugStatsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugStatsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *DebugHandler) Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error {
	return h.DebugHandler.Trace(ctx, in, out)
}

func (h *DebugHandler) Trace(ctx context.Context, stream server.Stream) error {

	m := new(TraceRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DebugHandler.Trace(ctx, m, &debugTraceStream{stream})

}

type Debug_TraceStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type debugTraceStream struct {
	stream server.Stream
}

func (x *debugTraceStream) Close() error {
	return x.stream.Close()
}

func (x *debugTraceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *debugTraceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugTraceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}
