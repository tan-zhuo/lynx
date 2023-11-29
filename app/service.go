package app

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
)

type ControlPlane interface {
	Limiter
	Registry
	Router
	Config
}

type Limiter interface {
	HttpRateLimit() middleware.Middleware
	GrpcRateLimit() middleware.Middleware
}

type Registry interface {
	NewServiceRegistry() registry.Registrar
	NewServiceDiscovery() registry.Discovery
}

type Router interface {
	NewNodeRouter(name string) selector.NodeFilter
}

type Config interface {
	Config(fileName string, group string) (config.Source, error)
}

func ServiceRegistry() registry.Registrar {
	return Lynx().ControlPlane().NewServiceRegistry()
}

func ServiceDiscovery() registry.Discovery {
	return Lynx().ControlPlane().NewServiceDiscovery()
}

func (a *LynxApp) ControlPlane() ControlPlane {
	return Lynx().cp
}
