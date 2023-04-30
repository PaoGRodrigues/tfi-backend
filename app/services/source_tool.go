package services

type NtopNG struct {
	UrlClient   string
	InterfaceId int
	Usr         string
	Pass        string
}

func NewTool(urlClient string, interfaceId int, usr string, pass string) *NtopNG {
	return &NtopNG{
		UrlClient:   urlClient,
		InterfaceId: interfaceId,
		Usr:         usr,
		Pass:        pass,
	}
}
