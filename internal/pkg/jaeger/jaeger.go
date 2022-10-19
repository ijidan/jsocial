package jaeger

import (
	"fmt"
	PkgConf "github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func NewJaeger(cf PkgConf.Jaeger, serviceName string) (opentracing.Tracer, io.Closer, error) {
	localAgentHostPort := fmt.Sprintf("%s:%d", cf.Host, cf.Port)
	cfg := &config.Configuration{
		ServiceName: serviceName,
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
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
