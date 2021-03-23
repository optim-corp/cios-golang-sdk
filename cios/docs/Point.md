# Point

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Location** | [**Location**](Location.md) |  | 
**Altitude** | **float32** |  | 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**ResourceOwnerId** | **string** |  | 

## Methods

### NewPoint

`func NewPoint(id string, name string, location Location, altitude float32, resourceOwnerId string, ) *Point`

NewPoint instantiates a new Point object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointWithDefaults

`func NewPointWithDefaults() *Point`

NewPointWithDefaults instantiates a new Point object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Point) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Point) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Point) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Point) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Point) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Point) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *Point) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Point) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Point) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetAltitude

`func (o *Point) GetAltitude() float32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *Point) GetAltitudeOk() (*float32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *Point) SetAltitude(v float32)`

SetAltitude sets Altitude field to given value.


### GetLabels

`func (o *Point) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Point) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Point) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Point) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *Point) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Point) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Point) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Point) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *Point) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *Point) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *Point) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *Point) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *Point) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Point) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Point) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


