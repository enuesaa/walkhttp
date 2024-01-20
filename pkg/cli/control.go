package cli

import (
	"log"
	"strings"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/spf13/cobra"
)

type SyncFileList struct {
	Items []SyncFile `json:"items"`
}
type SyncFile struct {
	IsDir bool `json:"isDir"`
	Path string `json:"path"`
}

func CreateControlCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "control",
		Short: "control",
		Run: func(cmd *cobra.Command, args []string) {
			host, _ := cmd.Flags().GetString("host")

			app := fiber.New()
			app.Use(logger.New())
			app.Get("/metrics", monitor.New())
			app.Get("/files", func(c *fiber.Ctx) error {
				res := SyncFileList {
					Items: make([]SyncFile, 0),
				}
				dirs, err := repos.Fs.ListDirs("testdata")
				if err != nil {
					return err
				}
				for _, dir := range dirs {
					res.Items = append(res.Items, SyncFile{
						IsDir: true,
						Path: dir,
					})
				}
				files, err := repos.Fs.ListFiles("testdata")
				if err != nil {
					return err
				}
				for _, file := range files {
					res.Items = append(res.Items, SyncFile{
						IsDir: false,
						Path: file,
					})
				}
				return c.JSON(res)
			})
			app.Get("/files/*", func(c *fiber.Ctx) error {
				path := c.Path()
				requestedPath := strings.TrimPrefix(path, "/files/")
				isDir, err := repos.Fs.IsDir(requestedPath)
				if err != nil {
					return err
				}
				if isDir {
					return nil
				}
				fbytes, err := repos.Fs.Read(requestedPath)
				if err != nil {
					return err
				}
				return c.SendString(string(fbytes))
			})
			if err := app.Listen(host); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().String("host", "localhost:3001", "host")

	return cmd
}