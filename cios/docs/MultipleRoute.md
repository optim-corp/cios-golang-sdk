# MultipleRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 

## Methods

### NewMultipleRoute

`func NewMultipleRoute() *MultipleRoute`

NewMultipleRoute instantiates a new MultipleRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleRouteWithDefaults

`func NewMultipleRouteWithDefaults() *MultipleRoute`

NewMultipleRouteWithDefaults instantiates a new MultipleRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleRoute) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleRoute) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleRoute) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MultipleRoute) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRoutes

`func (o *MultipleRoute) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *MultipleRoute) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *MultipleRoute) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *MultipleRoute) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


