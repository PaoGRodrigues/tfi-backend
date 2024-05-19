package repository

import "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"

type FakeSQLClient struct {
}

func NewFakeSQLClient() *FakeSQLClient {
	return &FakeSQLClient{}
}

func (client *FakeSQLClient) AddActiveFlows(flows []domains.ActiveFlow) error {
	return nil
}

func (client *FakeSQLClient) GetServerByAttr(attr string) (domains.Server, error) {

	server := domains.Server{}
	return server, nil
}

func (client *FakeSQLClient) GetClients() ([]domains.Client, error) {

	clients := []domains.Client{}
	return clients, nil
}

func (client *FakeSQLClient) GetServers() ([]domains.Server, error) {

	servers := []domains.Server{}
	return servers, nil
}

func (client *FakeSQLClient) GetFlowByKey(key string) (domains.ActiveFlow, error) {

	flow := domains.ActiveFlow{}
	return flow, nil
}
