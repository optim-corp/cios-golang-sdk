package ciossdk

import (
	_nethttp "net/http"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	cnv "github.com/fcfcqloow/go-advance/convert"
	xmath "github.com/fcfcqloow/go-advance/math"
	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/util"
)

func MakeGetModelsOpts() cios.ApiGetDeviceModelsRequest {
	return cios.ApiGetDeviceModelsRequest{}
}
func MakeGetEntitiesOpts() cios.ApiGetDeviceEntitiesRequest {
	return cios.ApiGetDeviceEntitiesRequest{}
}

func MakeGetLifecyclesOpts() cios.ApiGetDeviceEntitiesLifecyclesRequest {
	return cios.ApiGetDeviceEntitiesLifecyclesRequest{}
}

func (self *DeviceAssetManagement) GetModels(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) (response cios.MultipleDeviceModel, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleDeviceModel{}, nil, err
	}
	params.ApiService = self.ApiClient.DeviceAssetApi
	params.Ctx = self.withHost(ctx)
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_name = util.ToNil(params.P_name)
	params.P_componentKey = util.ToNil(params.P_componentKey)
	params.P_endEventAt = util.ToNil(params.P_endEventAt)
	params.P_componentValue = util.ToNil(params.P_componentValue)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	params.P_version = util.ToNil(params.P_version)
	return params.Execute()
}
func (self *DeviceAssetManagement) GetModelsAll(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error) {
	var (
		result      []cios.DeviceModel
		httpRes     *_nethttp.Response
		err         error
		_limit      = int64(1000)
		offset      = int64(0)
		getFunction = func(offset int64) (cios.MultipleDeviceModel, *_nethttp.Response, error) {
			return self.GetModels(ctx, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Models...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Models...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Models...)
		}
	}
	return result, httpRes, err
}
func (self *DeviceAssetManagement) GetModelsUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) ([]cios.DeviceModel, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetModelsAll(ctx, params)
}
func (self *DeviceAssetManagement) GetModelsMapByID(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceModelsRequest) (map[string]cios.DeviceModel, error) {
	models, _, err := self.GetModelsUnlimited(ctx, params)
	if err != nil {
		return nil, err
	}
	m := map[string]cios.DeviceModel{}
	for _, deviceModel := range models {
		m[deviceModel.Id] = deviceModel
	}
	return m, nil
}
func (self *DeviceAssetManagement) GetModel(ctx ciosctx.RequestCtx, name string) (cios.DeviceModel, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DeviceModel{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceAssetApi.GetDeviceModel(self.withHost(ctx), name).Execute()
	if err != nil {
		return cios.DeviceModel{}, httpResponse, err
	}
	return response.Model, httpResponse, err
}
func (self *DeviceAssetManagement) CreateModel(ctx ciosctx.RequestCtx, body cios.DeviceModelRequest) (cios.DeviceModel, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DeviceModel{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceAssetApi.CreateDeviceModel(self.withHost(ctx)).DeviceModelRequest(body).Execute()
	if err != nil {
		return cios.DeviceModel{}, httpResponse, err
	}
	return response.Model, httpResponse, err
}
func (self *DeviceAssetManagement) DeleteModel(ctx ciosctx.RequestCtx, name string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.DeviceAssetApi.DeleteDeviceModel(self.withHost(ctx), name).Execute()
}

func (self *DeviceAssetManagement) GetEntities(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) (response cios.MultipleDeviceModelEntity, httpResponse *_nethttp.Response, err error) {
	if err = self.refresh(); err != nil {
		return
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.DeviceAssetApi
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_key = util.ToNil(params.P_key)
	params.P_componentKey = util.ToNil(params.P_componentKey)
	params.P_serialNumber = util.ToNil(params.P_serialNumber)
	params.P_componentValue = util.ToNil(params.P_componentValue)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	params.P_deviceId = util.ToNil(params.P_deviceId)
	return params.Execute()
}
func (self *DeviceAssetManagement) GetEntitiesAll(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error) {
	var (
		result      []cios.DeviceModelsEntity
		httpRes     *_nethttp.Response
		err         error
		_limit      = int64(1000)
		offset      = int64(0)
		getFunction = func(offset int64) (cios.MultipleDeviceModelEntity, *_nethttp.Response, error) {
			return self.GetEntities(ctx, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Entities...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Entities...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Entities...)
		}
	}
	return result, httpRes, err
}
func (self *DeviceAssetManagement) GetEntitiesUnlimited(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) ([]cios.DeviceModelsEntity, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetEntitiesAll(ctx, params)
}
func (self *DeviceAssetManagement) GetEntitiesMapByID(ctx ciosctx.RequestCtx, params cios.ApiGetDeviceEntitiesRequest) (map[string]cios.DeviceModelsEntity, error) {
	devices, _, err := self.GetEntitiesUnlimited(ctx, params)
	if err != nil {
		return nil, err
	}
	m := map[string]cios.DeviceModelsEntity{}
	for _, entity := range devices {
		m[entity.Id] = entity
	}
	return m, nil
}
func (self *DeviceAssetManagement) GetEntity(ctx ciosctx.RequestCtx, key string) (cios.DeviceModelsEntity, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DeviceModelsEntity{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceAssetApi.GetDeviceEntity(self.withHost(ctx), key).Execute()
	if err != nil {
		return cios.DeviceModelsEntity{}, httpResponse, err
	}
	return response.Entity, httpResponse, err
}
func (self *DeviceAssetManagement) DeleteEntity(ctx ciosctx.RequestCtx, key string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.DeviceAssetApi.DeleteDeviceEntity(self.withHost(ctx), key).Execute()
}
func (self *DeviceAssetManagement) CreateEntity(ctx ciosctx.RequestCtx, name string, body cios.Inventory) (cios.DeviceModelsEntity, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.DeviceModelsEntity{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceAssetApi.CreateInventory(self.withHost(ctx), name).Inventory(body).Execute()
	if err != nil {
		return cios.DeviceModelsEntity{}, httpResponse, err
	}
	return response.Entity, httpResponse, err
}

func (self *DeviceAssetManagement) GetLifecycles(ctx ciosctx.RequestCtx, key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest) (response cios.MultipleLifeCycle, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleLifeCycle{}, nil, err
	}
	params.ApiService = self.ApiClient.DeviceAssetApi
	params.P_key = key
	params.Ctx = self.withHost(ctx)
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_afterId = util.ToNil(params.P_afterId)
	params.P_beforeId = util.ToNil(params.P_beforeId)
	params.P_componentId = util.ToNil(params.P_componentId)
	params.P_eventKind = util.ToNil(params.P_eventKind)
	params.P_eventMode = util.ToNil(params.P_eventMode)
	params.P_eventType = util.ToNil(params.P_eventType)
	params.P_endEventAt = util.ToNil(params.P_endEventAt)
	params.P_startEventAt = util.ToNil(params.P_startEventAt)
	return params.Execute()
}
func (self *DeviceAssetManagement) GetLifecyclesAll(ctx ciosctx.RequestCtx, key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error) {
	var (
		httpResponse *_nethttp.Response
		err          error
		result       []cios.LifeCycle
		_limit       = int64(1000)
		offset       = int64(0)
		getFunction  = func(offset int64) (cios.MultipleLifeCycle, *_nethttp.Response, error) {
			return self.GetLifecycles(ctx, key, params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+cnv.MustInt64(params.P_offset)))
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Lifecycles...)
			offset += 1000
			_limit -= 1000
			if _limit <= 0 {
				break
			}
		}
	} else {
		res, httpRes, err := getFunction(offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Lifecycles...)
		for offset = int64(1000); offset+cnv.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Lifecycles...)
		}
	}
	return result, httpResponse, err
}
func (self *DeviceAssetManagement) GetLifecyclesUnlimitedByEntities(ctx ciosctx.RequestCtx, entities []cios.DeviceModelsEntity, params cios.ApiGetDeviceEntitiesLifecyclesRequest) ([][]cios.LifeCycle, *_nethttp.Response, error) {
	var allLifecycles [][]cios.LifeCycle
	for _, modelEntity := range entities {
		lifecycles, httpResponse, err := self.GetLifecyclesUnlimited(ctx, modelEntity.Key, params)
		if err != nil {
			return nil, httpResponse, err
		}
		allLifecycles = append(allLifecycles, lifecycles)
	}
	return allLifecycles, nil, nil
}
func (self *DeviceAssetManagement) GetLifecyclesUnlimited(ctx ciosctx.RequestCtx, key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest) ([]cios.LifeCycle, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetLifecyclesAll(ctx, key, params)
}
func (self *DeviceAssetManagement) GetLifecycle(ctx ciosctx.RequestCtx, key, id string) (cios.LifeCycle, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.LifeCycle{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceAssetApi.GetDeviceEntitiesLifecycle(self.withHost(ctx), key, id).Execute()
	if err != nil {
		return cios.LifeCycle{}, nil, err
	}
	return response.Lifecycle, httpResponse, err
}
func (self *DeviceAssetManagement) CreateLifecycle(ctx ciosctx.RequestCtx, key string, body cios.LifeCycleRequest) (cios.LifeCycle, *_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return cios.LifeCycle{}, nil, err
	}
	response, httpResponse, err := self.ApiClient.DeviceAssetApi.CreateDeviceEntitiesLifecycle(self.withHost(ctx), key).LifeCycleRequest(body).Execute()
	if err != nil {
		return cios.LifeCycle{}, nil, err
	}
	return response.Lifecycle, httpResponse, err
}
func (self *DeviceAssetManagement) DeleteLifecycle(ctx ciosctx.RequestCtx, key string, id string) (*_nethttp.Response, error) {
	if err := self.refresh(); err != nil {
		return nil, err
	}
	return self.ApiClient.DeviceAssetApi.DeleteDeviceEntitiesLifecycle(self.withHost(ctx), key, id).Execute()
}

type DeviceEntitiesComponentRange struct {
	StartDeviceEntitiesComponent cios.DeviceEntitiesComponent
	EndDeviceEntitiesComponent   cios.DeviceEntitiesComponent
	StartEventAt                 string
	EndEventAt                   string
}

func ToBetweenLifecycle(lifecycles cios.LifeCycleStream) (result [][]DeviceEntitiesComponentRange) {
	streams := lifecycles.
		Sort(func(i, j int) bool { return lifecycles.Get(i).EventAt < lifecycles.Get(j).EventAt }).
		GroupBy(func(cycle cios.LifeCycle, _ int) string { return cnv.MustStr(cycle.AfterId) })
	for _, cycles := range streams {
		result = append(result, toBetweenLifecycle(cycles))
	}
	return
}
func toBetweenLifecycle(lifecycles cios.LifeCycleStream) (result []DeviceEntitiesComponentRange) {
	length := len(lifecycles)
	if length < 1 {
		return
	}
	result = append(result, toDeviceEntitiesComponentRange(nil, &lifecycles[0]))
	for i := 0; i < length-1; i++ {
		result = append(result, toDeviceEntitiesComponentRange(&lifecycles[i], &lifecycles[i+1]))
	}
	result = append(result, toDeviceEntitiesComponentRange(&lifecycles[length-1], nil))
	return
}
func toDeviceEntitiesComponentRange(start *cios.LifeCycle, end *cios.LifeCycle) DeviceEntitiesComponentRange {
	if start != nil && end == nil && start.AfterComponent.IsSet() {
		return DeviceEntitiesComponentRange{
			StartDeviceEntitiesComponent: *start.AfterComponent.Get(),
			StartEventAt:                 start.EventAt,
		}
	}
	if start != nil && end != nil && start.AfterComponent.IsSet() && end.BeforeComponent.IsSet() {
		return DeviceEntitiesComponentRange{
			StartDeviceEntitiesComponent: *start.AfterComponent.Get(),
			EndDeviceEntitiesComponent:   *end.BeforeComponent.Get(),
			StartEventAt:                 start.EventAt,
			EndEventAt:                   end.EventAt,
		}
	}
	if start == nil && end != nil && end.BeforeComponent.IsSet() {
		return DeviceEntitiesComponentRange{
			EndDeviceEntitiesComponent: *end.BeforeComponent.Get(),
			EndEventAt:                 end.EventAt,
		}
	}
	return DeviceEntitiesComponentRange{}
}
func CreateParentPartsMap(entities []cios.DeviceModelsEntity) map[string]string {
	var (
		dig            func(n cios.DeviceEntitiesComponent)
		parentPartsMap = map[string]string{}
	)
	dig = func(entity cios.DeviceEntitiesComponent) {
		if entity.Components != nil {
			for _, ent := range *entity.Components {
				if ent.DisplayInfo != nil {
					for _, display := range *ent.DisplayInfo {
						if display.Language == "en" {
							parentPartsMap[ent.Id] = display.Name
						}
					}
				}
				dig(ent)
			}
		}
	}
	cios.
		GenerateDeviceModelsEntityStream(entities).
		ForEach(func(device cios.DeviceModelsEntity, _ int) {
			if device.Components.IsSet() {
				dig(*device.Components.Get())
			}
		})
	return parentPartsMap
}
