package ciossdmock

import ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

type MockClient struct {
	ciossdk.CiosClient
	pubSub                ciossdk.PubSub
	account               ciossdk.Account
	deviceAssetManagement ciossdk.DeviceAssetManagement
	deviceManagement      ciossdk.DeviceManagement
	fileStorage           ciossdk.FileStorage
	geography             ciossdk.Geography
	auth                  ciossdk.Auth
	license               ciossdk.License
	contract              ciossdk.Contract
	video                 ciossdk.VideoStreaming
}

func (m *MockClient) PubSub() ciossdk.PubSub {
	return m.pubSub
}

func (m *MockClient) Account() ciossdk.Account {
	return m.account
}

func (m *MockClient) DeviceAssetManagement() ciossdk.DeviceAssetManagement {
	return m.deviceAssetManagement
}

func (m *MockClient) DeviceManagement() ciossdk.DeviceManagement {
	return m.deviceManagement
}

func (m *MockClient) FileStorage() ciossdk.FileStorage {
	return m.fileStorage
}

func (m *MockClient) Geography() ciossdk.Geography {
	return m.geography
}

func (m *MockClient) License() ciossdk.License {
	return m.license
}

func (m *MockClient) Contract() ciossdk.Contract {
	return m.contract
}

func (m *MockClient) Video() ciossdk.VideoStreaming {
	return m.video
}

func (m *MockClient) Auth() ciossdk.Auth {
	return m.auth
}

func (m *MockClient) SetPubSub(pubSub ciossdk.PubSub) {
	m.pubSub = pubSub
}

func (m *MockClient) SetAccount(account ciossdk.Account) {
	m.account = account
}

func (m *MockClient) SetDeviceAssetManagement(deviceAssetManagement ciossdk.DeviceAssetManagement) {
	m.deviceAssetManagement = deviceAssetManagement
}

func (m *MockClient) SetDeviceManagement(deviceManagement ciossdk.DeviceManagement) {
	m.deviceManagement = deviceManagement
}

func (m *MockClient) SetFileStorage(fileStorage ciossdk.FileStorage) {
	m.fileStorage = fileStorage
}

func (m *MockClient) SetGeography(geography ciossdk.Geography) {
	m.geography = geography
}

func (m *MockClient) SetAuth(auth ciossdk.Auth) {
	m.auth = auth
}

func (m *MockClient) SetLicense(license ciossdk.License) {
	m.license = license
}

func (m *MockClient) SetContract(contract ciossdk.Contract) {
	m.contract = contract
}

func (m *MockClient) SetVideo(video ciossdk.VideoStreaming) {
	m.video = video
}

func (m MockClient) String() string {
	panic("implement me")
}

func (m MockClient) Debug(debug bool) ciossdk.CiosClient {
	panic("implement me")
}

func (m MockClient) RequestScope(scope string) ciossdk.CiosClient {
	panic("implement me")
}

func (m MockClient) TokenExp() int64 {
	panic("implement me")
}

func (m MockClient) SetTokenExp(tokenExp int64) {
	panic("implement me")
}
