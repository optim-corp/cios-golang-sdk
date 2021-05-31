package ciossdmock

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/cios"
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"
)

type NoImplementDeviceAssetManagement struct{}

func (NoImplementDeviceAssetManagement) GetModels(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceModelsRequest) (cios.MultipleDeviceModel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetModelsAll(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetModelsUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetModelsMapByID(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceModelsRequest) (map[string]cios.DeviceModel, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetModel(ctx ciosctx.RequestCtx, s string) (cios.DeviceModel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) CreateModel(ctx ciosctx.RequestCtx, request cios.DeviceModelRequest) (cios.DeviceModel, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) DeleteModel(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetEntities(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceEntitiesRequest) (cios.MultipleDeviceModelEntity, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetEntitiesAll(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetEntitiesUnlimited(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetEntitiesMapByID(ctx ciosctx.RequestCtx, request cios.ApiGetDeviceEntitiesRequest) (map[string]cios.DeviceModelsEntity, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetEntity(ctx ciosctx.RequestCtx, s string) (cios.DeviceModelsEntity, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) DeleteEntity(ctx ciosctx.RequestCtx, s string) (*_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) CreateEntity(ctx ciosctx.RequestCtx, s string, inventory cios.Inventory) (cios.DeviceModelsEntity, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetLifecycles(ctx ciosctx.RequestCtx, s string, request cios.ApiGetDeviceEntitiesLifecyclesRequest) (cios.MultipleLifeCycle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetLifecyclesAll(ctx ciosctx.RequestCtx, s string, request cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetLifecyclesUnlimitedByEntities(ctx ciosctx.RequestCtx, entities []cios.DeviceModelsEntity, request cios.ApiGetDeviceEntitiesLifecyclesRequest) ([][]cios.LifeCycle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) GetLifecyclesUnlimited(ctx ciosctx.RequestCtx, s string, request cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) CreateLifecycle(ctx ciosctx.RequestCtx, s string, request cios.LifeCycleRequest) (cios.LifeCycle, *_nethttp.Response, error) {
	panic("implement me")
}

func (NoImplementDeviceAssetManagement) DeleteLifecycle(ctx ciosctx.RequestCtx, s string, s2 string) (*_nethttp.Response, error) {
	panic("implement me")
}
