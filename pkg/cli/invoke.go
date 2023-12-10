package cli

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	st "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func CreateInvokeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke",
		Run: func(cmd *cobra.Command, args []string) {
			url, _ := cmd.Flags().GetString("url")

			tp, err := NewTracerProvider()
			if err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
			tracer = otel.Tracer("aaaaa")
		
			hc := http.Client{
				Transport: otelhttp.NewTransport(http.DefaultTransport),
			}
			res, err := hc.Get(url)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
			defer res.Body.Close()
			resbodybytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
			fmt.Printf("%s\n", string(resbodybytes))

			if err := tp.ForceFlush(context.Background()); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
		},
	}
	cmd.Flags().String("url", "https://example.com", "invoke url")

	return cmd
}

func NewTracerProvider() (*st.TracerProvider, error) {
	endpoint := "http://localhost:14268/api/traces"
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
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

