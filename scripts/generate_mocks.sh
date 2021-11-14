#!/bin/bash

mockgen -destination tests/mocks/device/devices.go -source=app/device/domains/device.go DeviceGateway,DeviceRepository
