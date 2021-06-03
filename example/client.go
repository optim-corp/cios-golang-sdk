package example

import (
	"os"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
)

var client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
	AutoRefresh: true,
	Debug:       true,
	Urls: sdkmodel.CIOSUrl{
		LicenseUrl:               os.Getenv("CIOS_API_URL"),
		ContractUrl:              os.Getenv("CIOS_API_URL"),
		MessagingUrl:             os.Getenv("CIOS_API_URL"),
		LocationUrl:              os.Getenv("CIOS_API_URL"),
		AccountsUrl:              os.Getenv("CIOS_API_URL"),
		StorageUrl:               os.Getenv("CIOS_API_URL"),
		IamUrl:                   os.Getenv("CIOS_API_URL"),
		AuthUrl:                  os.Getenv("CIOS_Auth_URL"),
		VideoStreamingUrl:        os.Getenv("CIOS_API_URL"),
		DeviceManagementUrl:      os.Getenv("CIOS_API_URL"),
		DeviceMonitoringUrl:      os.Getenv("CIOS_API_URL"),
		DeviceAssetManagementUrl: os.Getenv("CIOS_API_URL"),
	},
	CustomClient: nil,
	AuthConfig: ciossdk.RefreshTokenAuth(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("REFRESH_TOKEN"),
		os.Getenv("SCOPE"),
	),
	WebSocketConfig: &ciossdk.WebSocketConfig{
		ReadTimeoutMilliSec:  10000,
		WriteTimeoutMilliSec: 10000,
	},
})
