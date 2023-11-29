package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-lynx/lynx/conf"
	gGrpc "google.golang.org/grpc"
)

type GrpcSubscribe struct {
	name string
	dis  registry.Discovery
	log  log.Logger
	tls  bool
}

type Option func(o *GrpcSubscribe)

// Subscribe subscribe GRPC service
func (g *GrpcSubscribe) Subscribe() *gGrpc.ClientConn {
	dfLog := log.NewHelper(g.log)
	endpoint := "discovery:///" + g.name
	con, err := grpc.Dial(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(g.dis),
		grpc.WithMiddleware(
			logging.Client(g.log),
			tracing.Client(),
		),
		grpc.WithTLSConfig(g.tlsLoad()),
	)
	if err != nil {
		dfLog.Error(err)
		panic(err)
	}
	return con
}

func (g *GrpcSubscribe) tlsLoad() *tls.Config {
	if !g.tls {
		return nil
	}

	s, err := Lynx().ControlPlane().Config("tls-root.yaml", g.name)

	c := config.New(
		config.WithSource(s),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var t conf.Tls
	if err := c.Scan(&t); err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM([]byte(t.RootCA)) {
		panic(err)
	}

	return &tls.Config{ServerName: g.name, RootCAs: certPool}
}

func WithServiceName(name string) Option {
	return func(o *GrpcSubscribe) {
		o.name = name
	}
}

func WithDiscovery(dis registry.Discovery) Option {
	return func(o *GrpcSubscribe) {
		o.dis = dis
	}
}

func WithLogger(log log.Logger) Option {
	return func(o *GrpcSubscribe) {
		o.log = log
	}
}

func EnableTls() Option {
	return func(o *GrpcSubscribe) {
		o.tls = true
	}
}

func NewGrpcSubscribe(opts ...Option) *GrpcSubscribe {
	gs := &GrpcSubscribe{
		tls: false,
	}
	for _, o := range opts {
		o(gs)
	}
	return gs
}
