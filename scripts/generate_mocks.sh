#!/bin/bash

mockgen -destination mocks/services/services.go -source=app/services/domains.go Tool,Terminal,NotificationChannel,Database
mockgen -destination mocks/ports/host/reader.go -source=app/ports/host/reader.go HostReader
mockgen -destination mocks/ports/host/db_repository.go -source=app/ports/host/db_repository.go HostDBRepository
mockgen -destination mocks/ports/host/blocker.go -source=app/ports/host/blocker.go HostBlocker
mockgen -destination mocks/ports/alert/reader.go -source=app/ports/alert/reader.go AlertReader
mockgen -destination mocks/ports/alert/notifier.go -source=app/ports/alert/notifier.go Notifier
mockgen -destination mocks/ports/notificationchannel/notificationChannel.go -source=app/ports/notificationchannel/channel.go NotificationChannel
mockgen -destination mocks/ports/traffic/reader.go -source=app/ports/traffic/reader.go TrafficReader
mockgen -destination mocks/ports/traffic/db_repository.go -source=app/ports/traffic/db_repository.go TrafficDBRepository