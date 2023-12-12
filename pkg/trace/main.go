package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	st "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
)

func NewTracerProvider() (*st.TracerProvider, error) {
	endpoint := "http://localhost:14268/api/traces"
	exporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}

	rsrc := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("ddd"),
	)
	tp := st.NewTracerProvider(
		st.WithBatcher(exporter),
		st.WithResource(rsrc),
		st.WithSampler(st.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
