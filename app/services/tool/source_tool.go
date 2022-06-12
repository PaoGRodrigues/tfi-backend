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

type DatabaseConn struct {
	Usr              string
	Pass             string
	ConnectionString string
}

func NewDatabase(Usr string, Pass string, ConnectionString string) *DatabaseConn {
	return &DatabaseConn{
		Usr:              Usr,
		Pass:             Pass,
		ConnectionString: ConnectionString,
	}
}
