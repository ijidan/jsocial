package jaeger

import (
	"fmt"
	"github.com/google/wire"
	PkgConf "github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

type ServiceName string

type Jaeger struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

func NewJaeger(conf *PkgConf.Jaeger, app *PkgConf.App) (*Jaeger, error) {
	localAgentHostPort := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	cfg := &config.Configuration{
		ServiceName: app.Name,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: localAgentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	instance := &Jaeger{
		Tracer: tracer,
		Closer: closer,
	}
	return instance, nil
}

var Provider = wire.NewSet(NewJaeger)
