package repository

type ActiveFlowsClient struct {
	urlClient   string
	interfaceId int
	usr         string
	pass        string
}

func NewActiveFlowClient(urlClient string, interfaceId int, usr string, pass string) *ActiveFlowsClient {
	return &ActiveFlowsClient{
		urlClient:   urlClient,
		interfaceId: interfaceId,
		usr:         usr,
		pass:        pass,
	}
}
