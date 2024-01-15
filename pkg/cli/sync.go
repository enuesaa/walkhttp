package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func CreateSyncCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "experimental sync command",
		Run: func(cmd *cobra.Command, args []string) {
			isServe, _ := cmd.Flags().GetBool("serve")
			isClient, _ := cmd.Flags().GetBool("client")
			host, _ := cmd.Flags().GetString("host")

			if isServe {
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
				return
			}

			if isClient {
				listres, err := http.Get(fmt.Sprintf("http://%s/files", host))
				if err != nil {
					log.Fatalf("Error: %s", err.Error())
				}
				defer listres.Body.Close()
				fmt.Printf("status(%s): %s", listres.Request.URL, listres.Status)
				fbytes, err := io.ReadAll(listres.Body)
				var files SyncFileList
				if err := json.Unmarshal(fbytes, &files); err != nil {
					log.Fatalf("Error: %s", err.Error())
				}

				for _, file := range files.Items {
					if file.IsDir {
						if err := repos.Fs.CreateDir(file.Path); err != nil {
							log.Fatalf("Error: %s", err.Error())
						}
					} else {
						res, err := http.Get(fmt.Sprintf("http://%s/files/%s", host, file.Path))
						if err != nil {
							log.Fatalf("Error: %s", err.Error())
						}
						defer res.Body.Close()
						fmt.Printf("status(%s): %s\n", res.Request.URL, res.Status)
						fbytes, err := io.ReadAll(res.Body)
						if err != nil {
							log.Fatalf("Error: %s", err.Error())
						}
						if err := repos.Fs.Create(file.Path, fbytes); err != nil {
							log.Fatalf("Error: %s", err.Error())
						}
					}
				}
				return
			}
		},
	}
	cmd.Flags().Bool("serve", false, "serve")
	cmd.Flags().Bool("client", false, "client")
	cmd.Flags().String("host", "localhost:3001", "host")

	return cmd
}
