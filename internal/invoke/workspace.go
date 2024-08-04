package invoke

func NewConfig() Workspace {
	return Workspace{
		BaseUrl: "https://",
	}
}

type Workspace struct {
	BaseUrl string `json:"baseUrl"` // like `https://example.com`
	// Entries []Entry `json:"entries"`
}

type Entry struct {
	Request EntryRequest `json:"request"`
	Response EntryResponse `json:"response"`
}

type EntryRequest struct {
	Method  string `json:"method"`
	Url     string `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    string `json:"body"`
	Started int    `json:"started"`
}

type EntryResponse struct {
	Status  int `json:"status"`
	Headers map[string][]string `json:"headers"`
	Body    string `json:"body"`
	Received int `json:"received"`
}
