package cli

import (
	"log"

	"github.com/enuesaa/walkin/controlweb"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/cobra"
)

type SyncFileList struct {
	Items []SyncFile `json:"items"`
}
type SyncFile struct {
	IsDir bool `json:"isDir"`
	Path string `json:"path"`
}

func CreateCtlCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ctl",
		Short: "serve admin console",
		Run: func(cmd *cobra.Command, args []string) {
			host, _ := cmd.Flags().GetString("host")

			app := fiber.New()
			app.Use(logger.New())
			app.Use(requestid.New())
			app.Get("/metrics", monitor.New())

			serveCtl := usecase.NewServeCtl()
			app.Post("/api/invoke", serveCtl.CreateInvoke)
			app.Get("/*", controlweb.Serve)

			if err := app.Listen(host); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().String("host", "localhost:3000", "host")

	return cmd
}