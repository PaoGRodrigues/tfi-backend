#!/bin/bash

mockgen -destination mocks/hosts/host.go -source=app/hosts/domain/host/host.go HostUseCase,HostService,HostsFilter,HostsStorage,HostsRepository
mockgen -destination mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficStorage,TrafficRepository,TrafficBytesParser,TrafficService
mockgen -destination mocks/alerts/alert.go -source=app/alerts/domains/alert.go AlertUseCase,AlertService,Notifier,AlertsSender
mockgen -destination mocks/services/services.go -source=app/services/domains.go Tool,Terminal,NotificationChannel,Database