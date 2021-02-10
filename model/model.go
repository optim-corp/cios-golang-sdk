package sdkmodel

import (
	"context"
)

const (
	REFRESH_TYPE AuthType = "refreshToken"
	CLIENT_TYPE  AuthType = "client"
	DEVICE_TYPE  AuthType = "device"
	NONE_TYPE    AuthType = "none"
	ALSTROEMERIA Stage    = "ALSTROEMERIA"
	VIOLA        Stage    = "VIOLA"
)
const (
	AUTH_INDEX = iota
	ACCOUNT_INDEX
	FILE_STORAGE_INDEX
	MESSAGING_INDEX
	DEVICE_INDEX
	DEVICE_ASSET_INDEX
	CONTRACT_INDEX
	LICENSE_INDEX
	LOCATION_INDEX
)

type (
	AuthType     = string
	AccessToken  = string
	ClientID     = string
	ClientSecret = string
	RefreshToken = string
	Scope        = string
	Assertion    = string
	TokenType    = string
	Stage        = string
	ExpiresIn    = int
	RequestCtx   context.Context
	CIOSUrl      struct {
		LicenseUrl               string
		ContractUrl              string
		MessagingUrl             string
		LocationUrl              string
		AccountsUrl              string
		StorageUrl               string
		IamUrl                   string
		AuthUrl                  string
		VideoStreamingUrl        string
		DeviceManagementUrl      string
		DeviceMonitoringUrl      string
		DeviceAssetManagementUrl string
	}
)

func ProdUrls() CIOSUrl {
	return CIOSUrl{
		// TODO: Add Url
		//LicenseUrl:               "https://device-management.optimcloudapis.com",
		//ContractUrl:              "https://device-management.optimcloudapis.com",
		MessagingUrl:             "https://messaging.optimcloudapis.com",
		LocationUrl:              "https://location.optimcloudapis.com",
		AccountsUrl:              "https://accounts.optimcloudapis.com",
		StorageUrl:               "https://storage.optimcloudapis.com",
		IamUrl:                   "https://iam.optimcloudapis.com",
		AuthUrl:                  "https://auth.optim.cloud",
		VideoStreamingUrl:        "https://video-streams.optimcloudapis.com",
		DeviceManagementUrl:      "https://device-management.optimcloudapis.com",
		DeviceMonitoringUrl:      "https://monitoring.optimcloudapis.com",
		DeviceAssetManagementUrl: "https://device-asset-lifecycle.optimcloudapis.com",
	}
}

func PreUrls() CIOSUrl {
	return CIOSUrl{
		// TODO: Add Url
		//LicenseUrl:               "https://device-management.optimcloudapis.com",
		//ContractUrl:              "https://device-management.optimcloudapis.com",
		MessagingUrl:             "https://messaging.preapis.cios.dev",
		LocationUrl:              "https://location.preapis.cios.dev",
		AccountsUrl:              "https://accounts.preapis.cios.dev",
		StorageUrl:               "https://storage.preapis.cios.dev",
		IamUrl:                   "https://iam.preapis.cios.dev",
		AuthUrl:                  "https://auth.pre.cios.dev",
		VideoStreamingUrl:        "https://video-streams.preapis.cios.dev",
		DeviceManagementUrl:      "https://device-management.preapis.cios.dev",
		DeviceMonitoringUrl:      "https://monitoring.preapis.cios.dev",
		DeviceAssetManagementUrl: "https://device-asset-lifecycle.preapis.cios.dev",
	}
}
