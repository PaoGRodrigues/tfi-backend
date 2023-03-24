#!/bin/bash

mockgen -destination mocks/hosts/host.go -source=app/hosts/domains/host.go HostUseCase,HostService,HostsFilter
mockgen -destination mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficRepoClient,TrafficRepository,TrafficActiveFlowsSearcher,ActiveFlowsStorage,TrafficService
mockgen -destination mocks/alerts/alert.go -source=app/alerts/domains/alert.go AlertUseCase,AlertService,Notifier,AlertsSender
mockgen -destination mocks/services/services.go -source=app/services/channel.go NotificationChannel