#!/bin/bash

mockgen -destination tests/mocks/sourceMock.go -package mocks github.com/PaoGRodrigues/tfi-backend/app/domains DeviceUseCase,DeviceRepository
