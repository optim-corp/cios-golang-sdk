# MultipleLifeCycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Lifecycles** | [**[]LifeCycle**](LifeCycle.md) |  | 

## Methods

### NewMultipleLifeCycle

`func NewMultipleLifeCycle(total int64, lifecycles []LifeCycle, ) *MultipleLifeCycle`

NewMultipleLifeCycle instantiates a new MultipleLifeCycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleLifeCycleWithDefaults

`func NewMultipleLifeCycleWithDefaults() *MultipleLifeCycle`

NewMultipleLifeCycleWithDefaults instantiates a new MultipleLifeCycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleLifeCycle) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleLifeCycle) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleLifeCycle) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetLifecycles

`func (o *MultipleLifeCycle) GetLifecycles() []LifeCycle`

GetLifecycles returns the Lifecycles field if non-nil, zero value otherwise.

### GetLifecyclesOk

`func (o *MultipleLifeCycle) GetLifecyclesOk() (*[]LifeCycle, bool)`

GetLifecyclesOk returns a tuple with the Lifecycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycles

`func (o *MultipleLifeCycle) SetLifecycles(v []LifeCycle)`

SetLifecycles sets Lifecycles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


