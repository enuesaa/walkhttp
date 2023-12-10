package cli

import (
	"log"
	"time"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

func CreateInvokeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke",
		Run: func(cmd *cobra.Command, args []string) {	
			// url, _ := cmd.Flags().GetString("url")

			// res, err := http.Get(url)
			// if err != nil {
			// 	log.Fatalf("Error: %s\n", err)
			// }
			// defer res.Body.Close()

			// resbodybytes, err := io.ReadAll(res.Body)
			// if err != nil {
			// 	log.Fatalf("Error: %s\n", err)
			// }
			// resbody := string(resbodybytes)
			// fmt.Printf("%s\n", resbody)

			// see https://qiita.com/ike_dai/items/f7d95852a86a46e1f19d

			cfg := jaegerConfig.Configuration{
				Sampler: &jaegerConfig.SamplerConfig{
					Type:  "const",
					Param: 1,
				},
				Reporter: &jaegerConfig.ReporterConfig{
					LogSpans:            true,
					BufferFlushInterval: 1 * time.Second,
					LocalAgentHostPort:  "127.0.0.1:5775",
				},
			}
			tracer, closer, err := cfg.New(
				"sample",
				jaegerConfig.Logger(jaeger.StdLogger),
			)
			if err != nil {
				log.Fatalf("Erraor: %s\n", err.Error())
			}
			opentracing.SetGlobalTracer(tracer)
			defer closer.Close()

			parent := opentracing.GlobalTracer().StartSpan("hello")
			defer parent.Finish()
			child := opentracing.GlobalTracer().StartSpan(
			"world", opentracing.ChildOf(parent.Context()))
			defer child.Finish()
		},
	}
	// cmd.Flags().String("url", "https://example.com", "invoke url")

	return cmd
}
