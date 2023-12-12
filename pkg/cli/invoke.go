package cli

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/trace"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer oteltrace.Tracer

func CreateInvokeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke",
		Run: func(cmd *cobra.Command, args []string) {
			url, _ := cmd.Flags().GetString("url")

			tp, err := trace.NewTracerProvider()
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
