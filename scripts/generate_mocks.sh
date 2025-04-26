#!/bin/bash

mockgen -destination mocks/hosts/host.go -source=app/domain/host/filters.go HostService,HostsStorage,HostsRepository
mockgen -destination mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficStorage,TrafficRepository,TrafficBytesParser,TrafficService
mockgen -destination mocks/alerts/alert.go -source=app/alerts/domains/alert.go AlertUseCase,AlertService,Notifier,AlertsSender
mockgen -destination mocks/services/services.go -source=app/services/domains.go Tool,Terminal,NotificationChannel,Database
mockgen -destination mocks/api/ports.go -source=app/api/ports.go GetLocalhostsUseCase
mockgen -destination mocks/ports/host/repository.go -source=app/ports/host/repository.go HostRepository
mockgen -destination mocks/ports/host/blocker.go -source=app/ports/host/blocker.go HostBlocker