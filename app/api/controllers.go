package api

import (
	alertUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/alert"
	hostUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/host"
	notificationChannelUseCases "github.com/PaoGRodrigues/tfi-backend/app/usecase/notificationchannel"
	trafficUsecases "github.com/PaoGRodrigues/tfi-backend/app/usecase/traffic"

	"github.com/gin-gonic/gin"
)

type Api struct {
	TrafficSearcher                      *trafficUsecases.GetTrafficFlowsUseCase
	GetLocalhostsUseCase                 *hostUsecases.GetLocalhostsUseCase
	GetTrafficFlowsPerDestinationUseCase *trafficUsecases.GetTrafficFlowsPerDestinationUseCase
	GetTrafficFlowsPerCountryUseCase     *trafficUsecases.GetTrafficFlowsPerCountryUseCase
	StoreTrafficFlowsUseCase             *trafficUsecases.StoreTrafficFlowsUseCase
	GetAlertsUseCase                     *alertUsecases.GetAlertsUseCase
	BlockHostUseCase                     *hostUsecases.BlockHostUseCase
	ConfigureNotificationChannelUseCase  *notificationChannelUseCases.ConfigureChannelUseCase
	NotifyAlertsUseCase                  *alertUsecases.NotifyAlertsUseCase
	StoreHostsUseCase                    *hostUsecases.StoreHostUseCase
	*gin.Engine
}
