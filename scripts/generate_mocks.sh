#!/bin/bash

mockgen -destination mocks/host/host.go -source=app/host/domains/host.go HostUseCase,HostRepository,LocalHostFilter
mockgen -destination mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficRepoClient,TrafficRepository,TrafficActiveFlowsSearcher,ActiveFlowsStorage