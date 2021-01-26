# Circle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Altitude** | Pointer to **float32** |  | [optional] 
**RadiusM** | Pointer to **float32** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**ResourceOwnerId** | Pointer to **string** |  | [optional] 

## Methods

### NewCircle

`func NewCircle() *Circle`

NewCircle instantiates a new Circle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircleWithDefaults

`func NewCircleWithDefaults() *Circle`

NewCircleWithDefaults instantiates a new Circle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Circle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Circle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Circle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Circle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Circle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Circle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Circle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Circle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAltitude

`func (o *Circle) GetAltitude() float32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *Circle) GetAltitudeOk() (*float32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *Circle) SetAltitude(v float32)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *Circle) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetRadiusM

`func (o *Circle) GetRadiusM() float32`

GetRadiusM returns the RadiusM field if non-nil, zero value otherwise.

### GetRadiusMOk

`func (o *Circle) GetRadiusMOk() (*float32, bool)`

GetRadiusMOk returns a tuple with the RadiusM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusM

`func (o *Circle) SetRadiusM(v float32)`

SetRadiusM sets RadiusM field to given value.

### HasRadiusM

`func (o *Circle) HasRadiusM() bool`

HasRadiusM returns a boolean if a field has been set.

### GetLocation

`func (o *Circle) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Circle) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Circle) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Circle) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLabels

`func (o *Circle) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Circle) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Circle) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Circle) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *Circle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Circle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Circle) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Circle) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *Circle) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *Circle) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *Circle) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *Circle) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *Circle) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Circle) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Circle) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.

### HasResourceOwnerId

`func (o *Circle) HasResourceOwnerId() bool`

HasResourceOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


