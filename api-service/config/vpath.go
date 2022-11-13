package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-micro.dev/v4/api/resolver"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewResolver(opts ...resolver.Option) resolver.Resolver {
	return &Resolver{opts: resolver.NewOptions(opts...)}
}

type Resolver struct {
	opts resolver.Options
}

// rpc请求
type RpcReq struct {
	Service string            `json:"service"`
	Method  string            `json:"method"`
	Request interface{}       `json:"request"`
}

func (r *Resolver) Resolve(req *http.Request) (*resolver.Endpoint, error) {
	if req.URL.Path == "/" {
		return nil, errors.New("unknown name")
	}
	parts := strings.Split(req.URL.Path[1:], "/")
	//  /rpc
	if len(parts) == 1 && parts[0] == "rpc" {
		rpcReq := RpcReq{}
		_ = json.NewDecoder(req.Body).Decode(&rpcReq)
		buf := bytes.Buffer{}
		err := json.NewEncoder(&buf).Encode(rpcReq.Request)
		if err == nil {
			_ = req.Body.Close()
			req.Body = ioutil.NopCloser(&buf)
		}
		return &resolver.Endpoint{
			Name:   rpcReq.Service,
			Host:   req.Host,
			Method: rpcReq.Method,
			Path:   req.URL.Path,
		}, nil
	}
	//  /:service/:method
	if len(parts) >= 2 {
		return &resolver.Endpoint{
			Name:   r.withNamespace(req, parts[0]),
			Host:   req.Host,
			Method: strings.Join(parts[1:], "."),
			Path:   req.URL.Path,
		}, nil
	}
	return nil, errors.New("unknown name")
}

func (r *Resolver) String() string {
	return "path"
}

func (r *Resolver) withNamespace(req *http.Request, parts ...string) string {
	ns := r.opts.Namespace(req)
	if len(ns) == 0 {
		return strings.Join(parts, ".")
	}
	return strings.Join(append([]string{ns}, parts...), ".")
}
