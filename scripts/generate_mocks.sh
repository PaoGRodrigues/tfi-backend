#!/bin/bash

mockgen -destination mocks/hosts/host.go -source=app/domain/host/models_temp.go HostsStorage,HostsRepository
mockgen -destination mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficStorage,TrafficRepository,TrafficBytesParser,TrafficService
mockgen -destination mocks/alerts/alert.go -source=app/alerts/domains/alert.go AlertUseCase,AlertService,Notifier,AlertsSender
mockgen -destination mocks/services/services.go -source=app/services/domains.go Tool,Terminal,NotificationChannel,Database
mockgen -destination mocks/ports/host/reader.go -source=app/ports/host/reader.go HostReader
mockgen -destination mocks/ports/host/db_repository.go -source=app/ports/host/db_repository.go HostDBRepository
mockgen -destination mocks/ports/host/blocker.go -source=app/ports/host/blocker.go HostBlocker