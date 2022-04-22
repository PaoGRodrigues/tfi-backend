#!/bin/bash

mockgen -destination mocks/device/devices.go -source=app/device/domains/device.go DeviceUseCase,DeviceRepository
mockgen -destination mocks/traffic/traffic.go -source=app/traffic/domains/traffic.go TrafficUseCase,TrafficRepository