# MultipleContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Contracts** | [**[]Contract**](Contract.md) |  | 

## Methods

### NewMultipleContract

`func NewMultipleContract(total int64, contracts []Contract, ) *MultipleContract`

NewMultipleContract instantiates a new MultipleContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleContractWithDefaults

`func NewMultipleContractWithDefaults() *MultipleContract`

NewMultipleContractWithDefaults instantiates a new MultipleContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleContract) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleContract) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleContract) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetContracts

`func (o *MultipleContract) GetContracts() []Contract`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *MultipleContract) GetContractsOk() (*[]Contract, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *MultipleContract) SetContracts(v []Contract)`

SetContracts sets Contracts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


