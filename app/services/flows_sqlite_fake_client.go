package services

import (
	"github.com/PaoGRodrigues/tfi-backend/app/domain/host"
	traffic_domains "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type FakeSQLClient struct {
}

func NewFakeSQLClient() *FakeSQLClient {
	return &FakeSQLClient{}
}

func (client *FakeSQLClient) AddActiveFlows(flows []traffic_domains.ActiveFlow) error {
	return nil
}

func (client *FakeSQLClient) GetServerByAttr(attr string) (traffic_domains.Server, error) {

	server := traffic_domains.Server{}
	return server, nil
}

func (client *FakeSQLClient) GetClients() ([]traffic_domains.Client, error) {

	clients := []traffic_domains.Client{}
	return clients, nil
}

func (client *FakeSQLClient) GetServers() ([]traffic_domains.Server, error) {

	servers := []traffic_domains.Server{}
	return servers, nil
}

func (client *FakeSQLClient) GetFlowByKey(key string) (traffic_domains.ActiveFlow, error) {

	flow := traffic_domains.ActiveFlow{}
	return flow, nil
}

func (client *FakeSQLClient) StoreHosts([]host.Host) error {
	return nil
}

func (client *FakeSQLClient) GetHostByIp(string) (host.Host, error) {
	return host.Host{}, nil
}
