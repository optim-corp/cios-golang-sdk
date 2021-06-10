# Contract

## Contract API

### interface

```
GetContracts(ciosctx.RequestCtx, cios.ApiGetContractsRequest) (cios.MultipleContract, *_nethttp.Response, error)
GetContractsAll(ciosctx.RequestCtx, cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error)
GetContractsUnlimited(ciosctx.RequestCtx, cios.ApiGetContractsRequest) ([]cios.Contract, *_nethttp.Response, error)
```
### Usage

#### Get Contracts max limit 1000

```go
options := srvcontract.MakeGetContractsOpts
contracts, httpReponse, err := client.Contract.GetContracts(ctx, options())
```


#### Get Contracts no limit

```go
options := srvcontract.MakeGetContractsOpts
contracts, httpReponse, err := client.Contract.GetContractsAll(ctx, options())
```


#### Get Contracts unlimited

```go
options := srvcontract.MakeGetContractsOpts
contracts, httpReponse, err := client.Contract.GetContractsUnlimited(ctx, options())
```