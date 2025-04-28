#!/bin/bash

mockgen -destination mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficStorage,TrafficRepository,TrafficBytesParser,TrafficService
mockgen -destination mocks/alerts/alert.go -source=app/domain/alert/model_temp.go Notifier,AlertsSender
mockgen -destination mocks/services/services.go -source=app/services/domains.go Tool,Terminal,NotificationChannel,Database
mockgen -destination mocks/ports/host/reader.go -source=app/ports/host/reader.go HostReader
mockgen -destination mocks/ports/host/db_repository.go -source=app/ports/host/db_repository.go HostDBRepository
mockgen -destination mocks/ports/host/blocker.go -source=app/ports/host/blocker.go HostBlocker
mockgen -destination mocks/ports/alert/reader.go -source=app/ports/alert/reader.go AlertReader