package web

type BehaviorType int

const (
	ReadLocalFiles BehaviorType = iota
	Proxy
)

type ProxyConfig struct {
	Url string
}

type Bahavior struct {
	Bahavior BehaviorType
	ProxyConfig ProxyConfig 
}

type ServeConfig struct {
	Paths map[string]Bahavior
}
