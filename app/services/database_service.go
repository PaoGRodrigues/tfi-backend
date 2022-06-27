package services

import "github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"

type Storage interface {
	CreateTables() error
	InsertActiveFlow(domains.ActiveFlow) (int, error)
	InsertClient(domains.Client, int) error
	InsertServer(domains.Server, int) error
	InsertProtocol(domains.Protocol, int) error
}

type DBService struct {
	Strg Storage
}

func NewDatabaseService(strg Storage) *DBService {
	return &DBService{
		Strg: strg,
	}
}
