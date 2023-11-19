package web

type BehaviorType int

const (
	ReadLocalFiles BehaviorType = iota
	Proxy
)

type ProxyConfig struct {
	Url string
}

type Behavior struct {
	Behavior BehaviorType
	ProxyConfig ProxyConfig 
}

type ServeConfig struct {
	Paths map[string]Behavior
}
