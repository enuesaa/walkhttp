package proxy

import (
	"fmt"
	"strings"

	"github.com/enuesaa/walkin/internal/web"
)

func ParseKVFlagValue(value string) (map[string]string, error) {
	kvs := map[string]string{}
	records := strings.Split(value, ",")

	for _, record := range records {
		kv := strings.Split(record, "=")
		if len(kv) != 2 {
			return kvs, fmt.Errorf("invalid flag value passed.")
		}

		key := kv[0]
		val := kv[1]
		kvs[key] = val
	}

	return kvs, nil
}

func ParseFlagsToServeConfig(readLocalFilesFlag []string, proxyFlag []string) (web.ServeConfig, error) {
	serveConfig := web.ServeConfig{
		Paths: map[string]web.Behavior{},
	}

	for _, value := range readLocalFilesFlag {
		localFilesKvs, err := ParseKVFlagValue(value)
		if err != nil {
			return serveConfig, err
		}
		path := ""
		for key, val := range localFilesKvs {
			switch key {
			case "path":
				path = val
			}
		} 
		serveConfig.Paths[path] = web.Behavior{
			Behavior: web.ReadLocalFiles,
		}
	}

	for _, value := range proxyFlag {
		proxyKvs, err := ParseKVFlagValue(value)
		if err != nil {
			return serveConfig, err
		}
		proxyPath := ""
		proxyUrl := ""
		for key, val := range proxyKvs {
			switch key {
			case "path":
				proxyPath = val
			case "url":
				proxyUrl = val
			}
		}

		serveConfig.Paths[proxyPath] = web.Behavior{
			Behavior: web.Proxy,
			ProxyConfig: web.ProxyConfig{
				Url: proxyUrl,
			},
		}
	}

	return serveConfig, nil
}
