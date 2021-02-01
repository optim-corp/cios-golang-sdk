# LifeCycleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventKind** | **string** |  | 
**EventMode** | **string** |  | 
**EventType** | **string** |  | 
**BeforeId** | Pointer to **string** |  | [optional] 
**BeforeComponent** | Pointer to [**NullableDeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | [optional] 
**AfterId** | Pointer to **string** |  | [optional] 
**AfterComponent** | Pointer to [**NullableDeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | [optional] 
**EventAt** | **string** | ナノ秒 | 
**Note** | Pointer to **string** |  | [optional] 
**ResourceOwnerId** | Pointer to **string** |  | [optional] 

## Methods

### NewLifeCycleRequest

`func NewLifeCycleRequest(eventKind string, eventMode string, eventType string, eventAt string, ) *LifeCycleRequest`

NewLifeCycleRequest instantiates a new LifeCycleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifeCycleRequestWithDefaults

`func NewLifeCycleRequestWithDefaults() *LifeCycleRequest`

NewLifeCycleRequestWithDefaults instantiates a new LifeCycleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventKind

`func (o *LifeCycleRequest) GetEventKind() string`

GetEventKind returns the EventKind field if non-nil, zero value otherwise.

### GetEventKindOk

`func (o *LifeCycleRequest) GetEventKindOk() (*string, bool)`

GetEventKindOk returns a tuple with the EventKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKind

`func (o *LifeCycleRequest) SetEventKind(v string)`

SetEventKind sets EventKind field to given value.


### GetEventMode

`func (o *LifeCycleRequest) GetEventMode() string`

GetEventMode returns the EventMode field if non-nil, zero value otherwise.

### GetEventModeOk

`func (o *LifeCycleRequest) GetEventModeOk() (*string, bool)`

GetEventModeOk returns a tuple with the EventMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMode

`func (o *LifeCycleRequest) SetEventMode(v string)`

SetEventMode sets EventMode field to given value.


### GetEventType

`func (o *LifeCycleRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *LifeCycleRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *LifeCycleRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetBeforeId

`func (o *LifeCycleRequest) GetBeforeId() string`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *LifeCycleRequest) GetBeforeIdOk() (*string, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *LifeCycleRequest) SetBeforeId(v string)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *LifeCycleRequest) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetBeforeComponent

`func (o *LifeCycleRequest) GetBeforeComponent() DeviceEntitiesComponent`

GetBeforeComponent returns the BeforeComponent field if non-nil, zero value otherwise.

### GetBeforeComponentOk

`func (o *LifeCycleRequest) GetBeforeComponentOk() (*DeviceEntitiesComponent, bool)`

GetBeforeComponentOk returns a tuple with the BeforeComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeComponent

`func (o *LifeCycleRequest) SetBeforeComponent(v DeviceEntitiesComponent)`

SetBeforeComponent sets BeforeComponent field to given value.

### HasBeforeComponent

`func (o *LifeCycleRequest) HasBeforeComponent() bool`

HasBeforeComponent returns a boolean if a field has been set.

### SetBeforeComponentNil

`func (o *LifeCycleRequest) SetBeforeComponentNil(b bool)`

 SetBeforeComponentNil sets the value for BeforeComponent to be an explicit nil

### UnsetBeforeComponent
`func (o *LifeCycleRequest) UnsetBeforeComponent()`

UnsetBeforeComponent ensures that no value is present for BeforeComponent, not even an explicit nil
### GetAfterId

`func (o *LifeCycleRequest) GetAfterId() string`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *LifeCycleRequest) GetAfterIdOk() (*string, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *LifeCycleRequest) SetAfterId(v string)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *LifeCycleRequest) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.

### GetAfterComponent

`func (o *LifeCycleRequest) GetAfterComponent() DeviceEntitiesComponent`

GetAfterComponent returns the AfterComponent field if non-nil, zero value otherwise.

### GetAfterComponentOk

`func (o *LifeCycleRequest) GetAfterComponentOk() (*DeviceEntitiesComponent, bool)`

GetAfterComponentOk returns a tuple with the AfterComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterComponent

`func (o *LifeCycleRequest) SetAfterComponent(v DeviceEntitiesComponent)`

SetAfterComponent sets AfterComponent field to given value.

### HasAfterComponent

`func (o *LifeCycleRequest) HasAfterComponent() bool`

HasAfterComponent returns a boolean if a field has been set.

### SetAfterComponentNil

`func (o *LifeCycleRequest) SetAfterComponentNil(b bool)`

 SetAfterComponentNil sets the value for AfterComponent to be an explicit nil

### UnsetAfterComponent
`func (o *LifeCycleRequest) UnsetAfterComponent()`

UnsetAfterComponent ensures that no value is present for AfterComponent, not even an explicit nil
### GetEventAt

`func (o *LifeCycleRequest) GetEventAt() string`

GetEventAt returns the EventAt field if non-nil, zero value otherwise.

### GetEventAtOk

`func (o *LifeCycleRequest) GetEventAtOk() (*string, bool)`

GetEventAtOk returns a tuple with the EventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAt

`func (o *LifeCycleRequest) SetEventAt(v string)`

SetEventAt sets EventAt field to given value.


### GetNote

`func (o *LifeCycleRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LifeCycleRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LifeCycleRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *LifeCycleRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *LifeCycleRequest) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *LifeCycleRequest) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *LifeCycleRequest) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.

### HasResourceOwnerId

`func (o *LifeCycleRequest) HasResourceOwnerId() bool`

HasResourceOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


