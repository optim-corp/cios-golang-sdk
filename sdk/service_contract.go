package ciossdk

import (
	_nethttp "net/http"

	"github.com/optim-corp/cios-golang-sdk/util"

	xmath "github.com/optim-kazuhiro-seida/go-advance-type/math"
	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/optim-corp/cios-golang-sdk/model"
)

type Contract struct {
	ApiClient *cios.APIClient
	Url       string
	refresh   func() (model.AccessToken, model.Scope, model.TokenType, model.ExpiresIn, error)
	autoR     bool
}

func MakeGetContractsOpts() cios.ApiGetContractsRequest {
	return cios.ApiGetContractsRequest{}
}

func (self Contract) GetContracts(params cios.ApiGetContractsRequest, ctx model.RequestCtx) (response cios.MultipleContract, httpResponse *_nethttp.Response, err error) {
	params.Ctx = ctx
	params.ApiService = self.ApiClient.ContractApi
	params.P_page = util.ToNil(params.P_page)
	response, httpResponse, err = params.Execute()
	if err != nil && self.autoR {
		if _, _, _, _, err = self.refresh(); err != nil {
			return
		}
		response, httpResponse, err = params.Execute()
	}
	return
}

func (self Contract) GetContractsAll(params cios.ApiGetContractsRequest, ctx model.RequestCtx) ([]cios.Contract, *_nethttp.Response, error) {
	var (
		result      []cios.Contract
		httpRes     *_nethttp.Response
		err         error
		offset      = int64(0)
		_limit      = int64(1000)
		getFunction = func(offset *int64) (cios.MultipleContract, *_nethttp.Response, error) {
			if offset != nil {
				params.P_offset = offset
			}
			tlimit := xmath.MinInt64(_limit, 1000)
			params.P_limit = &tlimit
			return self.GetContracts(params, ctx)
		}
	)
	if params.P_limit != nil {
		_limit = *params.P_limit
		for {
			res, httpRes, err := getFunction(&offset)
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
		res, httpRes, err := getFunction(&offset)
		if err != nil {
			return nil, httpRes, err
		}
		result = append(result, res.Contracts...)
		for offset = int64(1000); offset < res.Total; offset += 1000 {
			res, httpRes, err = getFunction(&offset)
			if err != nil {
				return nil, httpRes, err
			}
			result = append(result, res.Contracts...)
		}
	}
	return result, httpRes, err
}
func (self Contract) GetContractsUnlimited(params cios.ApiGetContractsRequest, ctx model.RequestCtx) ([]cios.Contract, *_nethttp.Response, error) {
	params.P_limit = nil
	return self.GetContractsAll(params, ctx)
}
