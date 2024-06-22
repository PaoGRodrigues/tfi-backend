package repository

import (
	"github.com/PaoGRodrigues/tfi-backend/app/services"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type FlowsRepo struct {
	Database services.Database
}

func NewFlowsRepo(database services.Database) *FlowsRepo {
	return &FlowsRepo{
		Database: database,
	}
}

func (fs *FlowsRepo) GetServerByAttr(attr string) (domains.Server, error) {
	flow, err := fs.Database.GetServerByAttr(attr)
	if err != nil {
		return domains.Server{}, err
	}
	return flow, nil
}

func (fs *FlowsRepo) GetClients() ([]domains.Client, error) {
	clients, err := fs.Database.GetClients()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (fs *FlowsRepo) GetServers() ([]domains.Server, error) {
	servers, err := fs.Database.GetServers()
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (fs *FlowsRepo) GetFlowByKey(key string) (domains.ActiveFlow, error) {
	flow, err := fs.Database.GetFlowByKey(key)
	if err != nil {
		return domains.ActiveFlow{}, err
	}
	return flow, nil
}

func (fs *FlowsRepo) StoreFlows(flows []domains.ActiveFlow) error {
	err := fs.Database.AddActiveFlows(flows)
	if err != nil {
		return err
	}
	return nil
}
