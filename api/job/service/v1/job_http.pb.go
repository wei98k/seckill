// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type JobHTTPServer interface {
	GetJob(context.Context, *GetJobRequest) (*GetJobReply, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterJobHTTPServer(s *http.Server, srv JobHTTPServer) {
	r := s.Route("/")
	r.GET("/job/{id}", _Job_GetJob0_HTTP_Handler(srv))
	r.GET("/helloworld/{name}", _Job_SayHello0_HTTP_Handler(srv))
}

func _Job_GetJob0_HTTP_Handler(srv JobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetJobRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.job.service.v1.Job/GetJob")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetJob(ctx, req.(*GetJobRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetJobReply)
		return ctx.Result(200, reply)
	}
}

func _Job_SayHello0_HTTP_Handler(srv JobHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.job.service.v1.Job/SayHello")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type JobHTTPClient interface {
	GetJob(ctx context.Context, req *GetJobRequest, opts ...http.CallOption) (rsp *GetJobReply, err error)
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type JobHTTPClientImpl struct {
	cc *http.Client
}

func NewJobHTTPClient(client *http.Client) JobHTTPClient {
	return &JobHTTPClientImpl{client}
}

func (c *JobHTTPClientImpl) GetJob(ctx context.Context, in *GetJobRequest, opts ...http.CallOption) (*GetJobReply, error) {
	var out GetJobReply
	pattern := "/job/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.job.service.v1.Job/GetJob"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *JobHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.job.service.v1.Job/SayHello"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
