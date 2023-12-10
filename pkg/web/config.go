package web

type BehaviorType int

const (
	ReadLocalFiles BehaviorType = iota
	Proxy
)

type ProxyConfig struct {
	Url string `json:"url"`
}

type Behavior struct {
	Behavior BehaviorType `json:"behavior"`
	ProxyConfig ProxyConfig `json:"proxyConfig"`
}

type ServeConfig struct {
	Paths map[string]Behavior `json:"paths"`
}
