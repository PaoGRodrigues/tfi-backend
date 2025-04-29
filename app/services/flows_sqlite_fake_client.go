package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	traffic "github.com/PaoGRodrigues/tfi-backend/app/domain/traffic"
)

type FakeSQLClient struct {
}

func NewFakeSQLClient() *FakeSQLClient {
	return &FakeSQLClient{}
}

func (client *FakeSQLClient) AddActiveFlows(flows []traffic.ActiveFlow) error {
	return nil
}

func (client *FakeSQLClient) GetServerByAttr(attr string) (traffic.Server, error) {

	server := traffic.Server{}
	return server, nil
}

func (client *FakeSQLClient) GetClients() ([]traffic.Client, error) {

	clients := []traffic.Client{}
	return clients, nil
}

func (client *FakeSQLClient) GetServers() ([]traffic.Server, error) {

	servers := []traffic.Server{}
	return servers, nil
}

func (client *FakeSQLClient) GetFlowByKey(key string) (traffic.ActiveFlow, error) {

	flow := traffic.ActiveFlow{}
	return flow, nil
}

func (client *FakeSQLClient) StoreHosts([]host.Host) error {
	return nil
}

func (client *FakeSQLClient) GetHostByIp(string) (host.Host, error) {
	return host.Host{}, nil
}
