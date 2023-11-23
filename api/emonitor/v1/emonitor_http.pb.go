// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.20.3
// source: emonitor/v1/emonitor.proto

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

const OperationEmonitorAddSite = "/api.emonitor.v1.Emonitor/AddSite"
const OperationEmonitorDeleteSite = "/api.emonitor.v1.Emonitor/DeleteSite"
const OperationEmonitorGetSite = "/api.emonitor.v1.Emonitor/GetSite"
const OperationEmonitorListSite = "/api.emonitor.v1.Emonitor/ListSite"
const OperationEmonitorUpdateSite = "/api.emonitor.v1.Emonitor/UpdateSite"

type EmonitorHTTPServer interface {
	AddSite(context.Context, *AddSiteRequest) (*AddSiteReply, error)
	DeleteSite(context.Context, *DeleteSiteRequest) (*DeleteSiteReply, error)
	GetSite(context.Context, *GetSiteRequest) (*GetSiteReply, error)
	ListSite(context.Context, *ListSiteRequest) (*ListSiteReply, error)
	UpdateSite(context.Context, *UpdateSiteRequest) (*UpdateSiteReply, error)
}

func RegisterEmonitorHTTPServer(s *http.Server, srv EmonitorHTTPServer) {
	r := s.Route("/")
	r.POST("/site/add", _Emonitor_AddSite0_HTTP_Handler(srv))
	r.POST("/site/update", _Emonitor_UpdateSite0_HTTP_Handler(srv))
	r.POST("/site/delete", _Emonitor_DeleteSite0_HTTP_Handler(srv))
	r.POST("/site/detail", _Emonitor_GetSite0_HTTP_Handler(srv))
	r.POST("/site/list", _Emonitor_ListSite0_HTTP_Handler(srv))
}

func _Emonitor_AddSite0_HTTP_Handler(srv EmonitorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddSiteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmonitorAddSite)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddSite(ctx, req.(*AddSiteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddSiteReply)
		return ctx.Result(200, reply)
	}
}

func _Emonitor_UpdateSite0_HTTP_Handler(srv EmonitorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSiteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmonitorUpdateSite)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSite(ctx, req.(*UpdateSiteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateSiteReply)
		return ctx.Result(200, reply)
	}
}

func _Emonitor_DeleteSite0_HTTP_Handler(srv EmonitorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSiteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmonitorDeleteSite)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSite(ctx, req.(*DeleteSiteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteSiteReply)
		return ctx.Result(200, reply)
	}
}

func _Emonitor_GetSite0_HTTP_Handler(srv EmonitorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSiteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmonitorGetSite)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSite(ctx, req.(*GetSiteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSiteReply)
		return ctx.Result(200, reply)
	}
}

func _Emonitor_ListSite0_HTTP_Handler(srv EmonitorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSiteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmonitorListSite)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSite(ctx, req.(*ListSiteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSiteReply)
		return ctx.Result(200, reply)
	}
}

type EmonitorHTTPClient interface {
	AddSite(ctx context.Context, req *AddSiteRequest, opts ...http.CallOption) (rsp *AddSiteReply, err error)
	DeleteSite(ctx context.Context, req *DeleteSiteRequest, opts ...http.CallOption) (rsp *DeleteSiteReply, err error)
	GetSite(ctx context.Context, req *GetSiteRequest, opts ...http.CallOption) (rsp *GetSiteReply, err error)
	ListSite(ctx context.Context, req *ListSiteRequest, opts ...http.CallOption) (rsp *ListSiteReply, err error)
	UpdateSite(ctx context.Context, req *UpdateSiteRequest, opts ...http.CallOption) (rsp *UpdateSiteReply, err error)
}

type EmonitorHTTPClientImpl struct {
	cc *http.Client
}

func NewEmonitorHTTPClient(client *http.Client) EmonitorHTTPClient {
	return &EmonitorHTTPClientImpl{client}
}

func (c *EmonitorHTTPClientImpl) AddSite(ctx context.Context, in *AddSiteRequest, opts ...http.CallOption) (*AddSiteReply, error) {
	var out AddSiteReply
	pattern := "/site/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmonitorAddSite))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EmonitorHTTPClientImpl) DeleteSite(ctx context.Context, in *DeleteSiteRequest, opts ...http.CallOption) (*DeleteSiteReply, error) {
	var out DeleteSiteReply
	pattern := "/site/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmonitorDeleteSite))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EmonitorHTTPClientImpl) GetSite(ctx context.Context, in *GetSiteRequest, opts ...http.CallOption) (*GetSiteReply, error) {
	var out GetSiteReply
	pattern := "/site/detail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmonitorGetSite))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EmonitorHTTPClientImpl) ListSite(ctx context.Context, in *ListSiteRequest, opts ...http.CallOption) (*ListSiteReply, error) {
	var out ListSiteReply
	pattern := "/site/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmonitorListSite))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EmonitorHTTPClientImpl) UpdateSite(ctx context.Context, in *UpdateSiteRequest, opts ...http.CallOption) (*UpdateSiteReply, error) {
	var out UpdateSiteReply
	pattern := "/site/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmonitorUpdateSite))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
