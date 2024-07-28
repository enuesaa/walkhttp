package invoke

func NewConfig() Workspace {
	return Workspace{
		BaseUrl: "https://",
	}
}

type Workspace struct {
	BaseUrl string `json:"baseUrl"` // like `https://example.com`
	// Histories []Invocation `json:"histories"`
}
