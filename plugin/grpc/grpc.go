package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-lynx/lynx/app"
	"github.com/go-lynx/lynx/plugin"
	"github.com/go-lynx/lynx/plugin/grpc/conf"
)

var name = "grpc"

type ServiceGrpc struct {
	grpc   *grpc.Server
	conf   conf.Grpc
	weight int
	tls    bool
}

type Option func(g *ServiceGrpc)

func EnableTls() Option {
	return func(g *ServiceGrpc) {
		g.tls = true
	}
}

func Weight(w int) Option {
	return func(g *ServiceGrpc) {
		g.weight = w
	}
}

func (g *ServiceGrpc) Weight() int {
	return g.weight
}

func (g *ServiceGrpc) Name() string {
	return name
}

func (g *ServiceGrpc) Load(b config.Value) (plugin.Plugin, error) {
	err := b.Scan(&g.conf)
	if err != nil {
		return nil, err
	}

	app.Lynx().GetHelper().Infof("Initializing GRPC service")

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			tracing.Server(tracing.WithTracerName(app.Name())),
			logging.Server(app.Lynx().GetLogger()),
			validate.Validator(),
			// Recovery program after exception
			recovery.Recovery(
				recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
					return nil
				}),
			),
		),
	}

	if g.conf.Network != "" {
		opts = append(opts, grpc.Network(g.conf.Network))
	}
	if g.conf.Addr != "" {
		opts = append(opts, grpc.Address(g.conf.Addr))
	}
	if g.conf.Timeout != nil {
		opts = append(opts, grpc.Timeout(g.conf.Timeout.AsDuration()))
	}
	if app.Lynx().ControlPlane() != nil {
		opts = append(opts, grpc.Middleware(app.Lynx().ControlPlane().HttpRateLimit()))
	}

	if g.tls {
		cert, err := tls.X509KeyPair([]byte(app.Lynx().Tls().Crt), []byte(app.Lynx().Tls().Key))
		if err != nil {
			return nil, err
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM([]byte(app.Lynx().Tls().RootCA)) {
			return nil, err
		}

		opts = append(opts, grpc.TLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientCAs:    certPool,
			ServerName:   app.Name(),
			ClientAuth:   tls.RequireAndVerifyClientCert,
		}))
	}

	g.grpc = grpc.NewServer(opts...)
	app.Lynx().GetHelper().Infof("GRPC service successfully initialized")
	return g, nil
}

func (g *ServiceGrpc) Unload() error {
	if g.grpc == nil {
		return nil
	}
	if err := g.grpc.Stop(nil); err != nil {
		app.Lynx().GetHelper().Error(err)
	}
	app.Lynx().GetHelper().Info("message", "Closing the GRPC resources")
	return nil
}

func Grpc(opts ...Option) plugin.Plugin {
	s := &ServiceGrpc{
		tls:    false,
		weight: 500,
	}
	for _, option := range opts {
		option(s)
	}
	return s
}
