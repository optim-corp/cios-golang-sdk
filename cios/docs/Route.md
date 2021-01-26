# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to [**[]RouteContents**](RouteContents.md) |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**ResourceOwnerId** | Pointer to **string** |  | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Route) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Route) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Route) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Route) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Route) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Route) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContents

`func (o *Route) GetContents() []RouteContents`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *Route) GetContentsOk() (*[]RouteContents, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *Route) SetContents(v []RouteContents)`

SetContents sets Contents field to given value.

### HasContents

`func (o *Route) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetLabels

`func (o *Route) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Route) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Route) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Route) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *Route) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Route) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Route) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Route) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *Route) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *Route) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *Route) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *Route) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *Route) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Route) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Route) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.

### HasResourceOwnerId

`func (o *Route) HasResourceOwnerId() bool`

HasResourceOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


