#!/bin/bash

mockgen -destination tests/mocks/device/gateway.go -source=app/device/domains/device.go DeviceGateway, DeviceRepository
