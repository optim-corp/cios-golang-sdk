package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"

	"github.com/optim-corp/cios-golang-sdk/util"

	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	xmath "github.com/optim-corp/cios-golang-sdk/util/go_advance_type/math"
)

func MakeGetContractsOpts() cios.ApiGetContractsRequest {
	return cios.ApiGetContractsRequest{}
}

func (self *Contract) GetContracts(params cios.ApiGetContractsRequest, ctx sdkmodel.RequestCtx) (response cios.MultipleContract, httpResponse *_nethttp.Response, err error) {
	if err := self.refresh(); err != nil {
		return cios.MultipleContract{}, nil, err
	}
	params.Ctx = self.withHost(ctx)
	params.ApiService = self.ApiClient.ContractApi
	params.P_page = util.ToNil(params.P_page)
	return params.Execute()
}

func (self *Contract) GetContractsAll(params cios.ApiGetContractsRequest, ctx sdkmodel.RequestCtx) ([]cios.Contract, *_nethttp.Response, error) {
	var (
		result      []cios.Contract
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset int64) (cios.MultipleContract, *_nethttp.Response, error) {
			return self.GetContracts(params.Limit(xmath.MinInt64(_limit, 1000)).Offset(offset+convert.MustInt64(params.P_offset)), ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Contracts...)
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
		result = append(result, res.Contracts...)
		for offset = int64(1000); offset+convert.MustInt64(params.P_offset) < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Contracts...)
		}
	}
	return result, httpRes, err
}
func (self *Contract) GetContractsUnlimited(params cios.ApiGetContractsRequest, ctx sdkmodel.RequestCtx) ([]cios.Contract, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetContractsAll(params, ctx)
}
