package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"

	"github.com/optim-corp/cios-golang-sdk/cios"

	"github.com/optim-corp/cios-golang-sdk/model"
)

type DeviceAssetManagement struct {
	ApiClient *cios.APIClient
	Url       string
	refresh   func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error)
	autoR     bool
}

func MakeGetModelsOpts() cios.ApiGetDeviceModelsRequest {
	return cios.ApiGetDeviceModelsRequest{}
}
func MakeGetEntitiesOpts() cios.ApiGetDeviceEntitiesRequest {
	return cios.ApiGetDeviceEntitiesRequest{}
}

func MakeGetLifecyclesOpts() cios.ApiGetDeviceEntitiesLifecyclesRequest {
	return cios.ApiGetDeviceEntitiesLifecyclesRequest{}
}

func (self DeviceAssetManagement) GetModels(params cios.ApiGetDeviceModelsRequest, ctx model.RequestCtx) (response cios.MultipleDeviceModel, httpResponse *_nethttp.Response, err error) {
	params.ApiService = self.ApiClient.DeviceAssetApi
	params.Ctx = ctx
	response, httpResponse, err = params.Execute()
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_name = util.ToNil(params.P_name)
	params.P_componentKey = util.ToNil(params.P_componentKey)
	params.P_endEventAt = util.ToNil(params.P_endEventAt)
	params.P_componentValue = util.ToNil(params.P_componentValue)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	params.P_version = util.ToNil(params.P_version)
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}
func (self DeviceAssetManagement) GetModelsAll(params cios.ApiGetDeviceModelsRequest, ctx model.RequestCtx) ([]cios.DeviceModel, *_nethttp.Response, error) {
	var (
		result      []cios.DeviceModel
		httpRes     *_nethttp.Response
		err         error
		_limit      = int64(1000)
		offset      = int64(0)
		getFunction = func(offset *int64) (cios.MultipleDeviceModel, *_nethttp.Response, error) {
			if offset != nil {
				params.P_offset = offset
			}
			tlimit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &tlimit
			return self.GetModels(params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(&offset)
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
		res, httpRes, err := getFunction(&offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Models...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Models...)
		}
	}
	return result, httpRes, err
}
func (self DeviceAssetManagement) GetModelsUnlimited(params cios.ApiGetDeviceModelsRequest, ctx model.RequestCtx) ([]cios.DeviceModel, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetModelsAll(params, ctx)
}
func (self DeviceAssetManagement) GetModelsMapByID(ctx model.RequestCtx) (map[string]cios.DeviceModel, error) {
	models, _, err := self.GetModelsUnlimited(MakeGetModelsOpts(), ctx)
	if err != nil {
		return nil, err
	}
	m := map[string]cios.DeviceModel{}
	for _, deviceModel := range models {
		m[deviceModel.Id] = deviceModel
	}
	return m, nil
}
func (self DeviceAssetManagement) GetModel(name string, ctx model.RequestCtx) (cios.DeviceModel, *_nethttp.Response, error) {
	request := self.ApiClient.DeviceAssetApi.GetDeviceModel(ctx, name)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.DeviceModel{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.DeviceModel{}, httpResponse, err
		}
	}
	return response.Model, httpResponse, err
}
func (self DeviceAssetManagement) CreateModel(body cios.DeviceModelRequest, ctx model.RequestCtx) (cios.DeviceModel, *_nethttp.Response, error) {
	request := self.ApiClient.DeviceAssetApi.CreateDeviceModel(ctx).DeviceModelRequest(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.DeviceModel{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.DeviceModel{}, httpResponse, err
		}
	}
	return response.Model, httpResponse, err
}
func (self DeviceAssetManagement) DeleteModel(name string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	request := self.ApiClient.DeviceAssetApi.DeleteDeviceModel(ctx, name)
	httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return httpResponse, err
		}
		httpResponse, err = request.Execute()
	}
	return httpResponse, err
}

func (self DeviceAssetManagement) GetEntities(params cios.ApiGetDeviceEntitiesRequest, ctx model.RequestCtx) (response cios.MultipleDeviceModelEntity, httpResponse *_nethttp.Response, err error) {
	params.Ctx = ctx
	params.ApiService = self.ApiClient.DeviceAssetApi
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_key = util.ToNil(params.P_key)
	params.P_componentKey = util.ToNil(params.P_componentKey)
	params.P_serialNumber = util.ToNil(params.P_serialNumber)
	params.P_componentValue = util.ToNil(params.P_componentValue)
	params.P_resourceOwnerId = util.ToNil(params.P_resourceOwnerId)
	params.P_deviceId = util.ToNil(params.P_deviceId)
	response, httpResponse, err = params.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}
func (self DeviceAssetManagement) GetEntitiesAll(params cios.ApiGetDeviceEntitiesRequest, ctx model.RequestCtx) ([]cios.DeviceModelsEntity, *_nethttp.Response, error) {
	var (
		result      []cios.DeviceModelsEntity
		httpRes     *_nethttp.Response
		err         error
		_limit      = int64(1000)
		offset      = int64(0)
		getFunction = func(offset *int64) (cios.MultipleDeviceModelEntity, *_nethttp.Response, error) {
			if offset != nil {
				params.P_offset = offset
			}
			tlimit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &tlimit
			return self.GetEntities(params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(&offset)
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
		res, httpRes, err := getFunction(&offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Entities...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Entities...)
		}
	}
	return result, httpRes, err
}
func (self DeviceAssetManagement) GetEntitiesUnlimited(params cios.ApiGetDeviceEntitiesRequest, ctx model.RequestCtx) ([]cios.DeviceModelsEntity, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetEntitiesAll(params, ctx)
}
func (self DeviceAssetManagement) GetEntitiesMapByID(ctx model.RequestCtx) (map[string]cios.DeviceModelsEntity, error) {
	devices, _, err := self.GetEntitiesUnlimited(MakeGetEntitiesOpts(), ctx)
	if err != nil {
		return nil, err
	}
	m := map[string]cios.DeviceModelsEntity{}
	for _, entity := range devices {
		m[entity.Id] = entity
	}
	return m, nil
}
func (self DeviceAssetManagement) GetEntity(key string, ctx model.RequestCtx) (cios.DeviceModelsEntity, *_nethttp.Response, error) {
	request := self.ApiClient.DeviceAssetApi.GetDeviceEntity(ctx, key)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.DeviceModelsEntity{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.DeviceModelsEntity{}, httpResponse, err
		}
	}
	return response.Entity, httpResponse, err
}
func (self DeviceAssetManagement) DeleteEntity(key string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	request := self.ApiClient.DeviceAssetApi.DeleteDeviceEntity(ctx, key)
	herr, err := request.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return herr, err
		}
		return request.Execute()
	}
	return herr, err
}
func (self DeviceAssetManagement) CreateEntity(name string, body cios.Inventory, ctx model.RequestCtx) (cios.DeviceModelsEntity, *_nethttp.Response, error) {
	request := self.ApiClient.DeviceAssetApi.CreateInventory(ctx, name).Inventory(body)
	response, httpResponse, err := request.Execute()
	if err != nil {
		if self.autoR {
			if _, _, _, _, err = self.refresh(); err != nil {
				return cios.DeviceModelsEntity{}, httpResponse, err
			}
			response, httpResponse, err = request.Execute()
		}
		if err != nil {
			return cios.DeviceModelsEntity{}, httpResponse, err
		}
	}
	return response.Entity, httpResponse, err
}

func (self DeviceAssetManagement) GetLifecycles(key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest, ctx model.RequestCtx) (response cios.MultipleLifeCycle, httpResponse *_nethttp.Response, err error) {
	params.ApiService = self.ApiClient.DeviceAssetApi
	params.P_key = key
	params.Ctx = ctx
	params.P_order = util.ToNil(params.P_order)
	params.P_orderBy = util.ToNil(params.P_orderBy)
	params.P_afterId = util.ToNil(params.P_afterId)
	params.P_beforeId = util.ToNil(params.P_beforeId)
	params.P_componentId = util.ToNil(params.P_componentId)
	params.P_eventKind = util.ToNil(params.P_eventKind)
	params.P_eventMode = util.ToNil(params.P_eventMode)
	params.P_eventType = util.ToNil(params.P_eventType)
	params.P_endEventAt = util.ToNil(params.P_endEventAt)
	response, httpResponse, err = params.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}
func (self DeviceAssetManagement) GetLifecyclesAll(key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest, ctx model.RequestCtx) ([]cios.LifeCycle, *_nethttp.Response, error) {
	var (
		httpResponse *_nethttp.Response
		err          error
		result       []cios.LifeCycle
		_limit       = int64(1000)
		offset       = int64(0)
		getFunction  = func(offset *int64) (cios.MultipleLifeCycle, *_nethttp.Response, error) {
			if offset != nil {
				params.P_offset = offset
			}
			tlimit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &tlimit
			return self.GetLifecycles(key, params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(&offset)
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
		res, httpRes, err := getFunction(&offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Lifecycles...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Lifecycles...)
		}
	}
	return result, httpResponse, err
}
func (self DeviceAssetManagement) GetLifecyclesUnlimitedByEntities(entities []cios.DeviceModelsEntity, params cios.ApiGetDeviceEntitiesLifecyclesRequest, ctx model.RequestCtx) ([][]cios.LifeCycle, *_nethttp.Response, error) {
	var allLifecycles [][]cios.LifeCycle
	for _, modelEntity := range entities {
		lifecycles, httpResponse, err := self.GetLifecyclesUnlimited(modelEntity.Key, params, ctx)
		if err != nil {
			return nil, httpResponse, err
		}
		allLifecycles = append(allLifecycles, lifecycles)
	}
	return allLifecycles, nil, nil
}
func (self DeviceAssetManagement) GetLifecyclesUnlimited(key string, params cios.ApiGetDeviceEntitiesLifecyclesRequest, ctx model.RequestCtx) ([]cios.LifeCycle, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetLifecyclesAll(key, params, ctx)
}
func (self DeviceAssetManagement) DeleteLifecycle(key string, id string, ctx model.RequestCtx) (*_nethttp.Response, error) {
	request := self.ApiClient.DeviceAssetApi.DeleteDeviceEntitiesLifecycle(ctx, key, id)
	httpResponse, err := request.Execute()
	if err != nil && self.autoR {
		self.refresh()
		return request.Execute()
	}
	return httpResponse, err
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
		GroupBy(func(cycle cios.LifeCycle, _ int) string { return convert.MustStr(cycle.AfterId) })
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
	if start != nil && end == nil && start.AfterComponent != nil {
		return DeviceEntitiesComponentRange{
			StartDeviceEntitiesComponent: *start.AfterComponent,
			StartEventAt:                 start.EventAt,
		}
	}
	if start != nil && end != nil && start.AfterComponent != nil && end.BeforeComponent != nil {
		return DeviceEntitiesComponentRange{
			StartDeviceEntitiesComponent: *start.AfterComponent,
			EndDeviceEntitiesComponent:   *end.BeforeComponent,
			StartEventAt:                 start.EventAt,
			EndEventAt:                   end.EventAt,
		}
	}
	if start == nil && end != nil && end.BeforeComponent != nil {
		return DeviceEntitiesComponentRange{
			EndDeviceEntitiesComponent: *end.BeforeComponent,
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
		ForEach(func(device cios.DeviceModelsEntity, _ int) { dig(device.Components) })
	return parentPartsMap
}
