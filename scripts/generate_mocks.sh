#!/bin/bash

mockgen -destination tests/mocks/device/devices.go -source=app/device/domains/device.go DeviceUseCase,DeviceRepository
mockgen -destination tests/mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficRepository