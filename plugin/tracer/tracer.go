package tracer

import (
	"bytes"
	"github.com/go-lynx/lynx/boot"
	"github.com/go-lynx/lynx/conf"
	"github.com/go-lynx/lynx/plugin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var plugName = "tracer"

type PlugTracer struct {
	weight int
}

type Option func(t *PlugTracer)

func Weight(w int) Option {
	return func(t *PlugTracer) {
		t.weight = w
	}
}

func (t *PlugTracer) Weight() int {
	return t.weight
}

func (t *PlugTracer) Name() string {
	return plugName
}

func (t *PlugTracer) Load(b *conf.Bootstrap) (plugin.Plugin, error) {
	boot.GetHelper().Infof("Initializing link monitoring component")
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(b.Server.Tracer.Addr)))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.WriteString(b.Server.Name)
	buf.WriteString("-")
	buf.WriteString(b.Server.Version)
	tp := traceSdk.NewTracerProvider(
		traceSdk.WithSampler(traceSdk.ParentBased(traceSdk.TraceIDRatioBased(1.0))),
		traceSdk.WithBatcher(exp),
		traceSdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(buf.String()),
			attribute.String("exporter", "jaeger"),
			attribute.Float64("float", 312.23),
		)),
	)
	otel.SetTracerProvider(tp)
	boot.GetHelper().Infof("Link monitoring component successfully initialized")
	return t, nil
}

func (t *PlugTracer) Unload() error {
	return nil
}

func Tracer(opts ...Option) plugin.Plugin {
	t := &PlugTracer{
		weight: 700,
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}