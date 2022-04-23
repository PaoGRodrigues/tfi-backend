package domains

type Tool struct {
	UrlClient   string
	InterfaceId int
	Usr         string
	Pass        string
}

func NewTool(urlClient string, interfaceId int, usr string, pass string) *Tool {
	return &Tool{
		UrlClient:   urlClient,
		InterfaceId: interfaceId,
		Usr:         usr,
		Pass:        pass,
	}
}
