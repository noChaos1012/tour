package tracer

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,            //是否启用
			BufferFlushInterval: 1 * time.Second, //刷新频率
			LocalAgentHostPort:  agentHostPort,   //上报的地址
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return tracer, closer, err
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil

}
