package config

type Config struct {
	Paths []ConfigPath `json:"paths"`
}
type ConfigPath struct {
	Path string `json:"path"`
	Url string `json:"url"`
	OriginRequestHeaders map[string]string `json:"originRequestHeaders"`
}
