# PointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Altitude** | Pointer to **float32** |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**ResourceOwnerId** | **string** |  | 

## Methods

### NewPointRequest

`func NewPointRequest(resourceOwnerId string, ) *PointRequest`

NewPointRequest instantiates a new PointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointRequestWithDefaults

`func NewPointRequestWithDefaults() *PointRequest`

NewPointRequestWithDefaults instantiates a new PointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PointRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PointRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *PointRequest) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PointRequest) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PointRequest) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PointRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAltitude

`func (o *PointRequest) GetAltitude() float32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *PointRequest) GetAltitudeOk() (*float32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *PointRequest) SetAltitude(v float32)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *PointRequest) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetLabels

`func (o *PointRequest) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PointRequest) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PointRequest) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PointRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *PointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *PointRequest) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *PointRequest) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *PointRequest) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *PointRequest) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *PointRequest) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *PointRequest) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *PointRequest) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


