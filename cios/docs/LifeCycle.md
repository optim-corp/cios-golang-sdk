# LifeCycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EventKind** | **string** |  | 
**EventMode** | **string** |  | 
**EventType** | **string** |  | 
**BeforeId** | Pointer to **string** |  | [optional] 
**BeforeComponent** | Pointer to [**DeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | [optional] 
**AfterId** | Pointer to **string** |  | [optional] 
**AfterComponent** | Pointer to [**DeviceEntitiesComponent**](DeviceEntitiesComponent.md) |  | [optional] 
**EventAt** | **string** | ナノ秒 | 
**Note** | Pointer to **string** |  | [optional] 
**ResourceOwnerId** | **string** |  | 

## Methods

### NewLifeCycle

`func NewLifeCycle(id string, eventKind string, eventMode string, eventType string, eventAt string, resourceOwnerId string, ) *LifeCycle`

NewLifeCycle instantiates a new LifeCycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifeCycleWithDefaults

`func NewLifeCycleWithDefaults() *LifeCycle`

NewLifeCycleWithDefaults instantiates a new LifeCycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LifeCycle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LifeCycle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LifeCycle) SetId(v string)`

SetId sets Id field to given value.


### GetEventKind

`func (o *LifeCycle) GetEventKind() string`

GetEventKind returns the EventKind field if non-nil, zero value otherwise.

### GetEventKindOk

`func (o *LifeCycle) GetEventKindOk() (*string, bool)`

GetEventKindOk returns a tuple with the EventKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKind

`func (o *LifeCycle) SetEventKind(v string)`

SetEventKind sets EventKind field to given value.


### GetEventMode

`func (o *LifeCycle) GetEventMode() string`

GetEventMode returns the EventMode field if non-nil, zero value otherwise.

### GetEventModeOk

`func (o *LifeCycle) GetEventModeOk() (*string, bool)`

GetEventModeOk returns a tuple with the EventMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMode

`func (o *LifeCycle) SetEventMode(v string)`

SetEventMode sets EventMode field to given value.


### GetEventType

`func (o *LifeCycle) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *LifeCycle) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *LifeCycle) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetBeforeId

`func (o *LifeCycle) GetBeforeId() string`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *LifeCycle) GetBeforeIdOk() (*string, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *LifeCycle) SetBeforeId(v string)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *LifeCycle) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetBeforeComponent

`func (o *LifeCycle) GetBeforeComponent() DeviceEntitiesComponent`

GetBeforeComponent returns the BeforeComponent field if non-nil, zero value otherwise.

### GetBeforeComponentOk

`func (o *LifeCycle) GetBeforeComponentOk() (*DeviceEntitiesComponent, bool)`

GetBeforeComponentOk returns a tuple with the BeforeComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeComponent

`func (o *LifeCycle) SetBeforeComponent(v DeviceEntitiesComponent)`

SetBeforeComponent sets BeforeComponent field to given value.

### HasBeforeComponent

`func (o *LifeCycle) HasBeforeComponent() bool`

HasBeforeComponent returns a boolean if a field has been set.

### GetAfterId

`func (o *LifeCycle) GetAfterId() string`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *LifeCycle) GetAfterIdOk() (*string, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *LifeCycle) SetAfterId(v string)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *LifeCycle) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.

### GetAfterComponent

`func (o *LifeCycle) GetAfterComponent() DeviceEntitiesComponent`

GetAfterComponent returns the AfterComponent field if non-nil, zero value otherwise.

### GetAfterComponentOk

`func (o *LifeCycle) GetAfterComponentOk() (*DeviceEntitiesComponent, bool)`

GetAfterComponentOk returns a tuple with the AfterComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterComponent

`func (o *LifeCycle) SetAfterComponent(v DeviceEntitiesComponent)`

SetAfterComponent sets AfterComponent field to given value.

### HasAfterComponent

`func (o *LifeCycle) HasAfterComponent() bool`

HasAfterComponent returns a boolean if a field has been set.

### GetEventAt

`func (o *LifeCycle) GetEventAt() string`

GetEventAt returns the EventAt field if non-nil, zero value otherwise.

### GetEventAtOk

`func (o *LifeCycle) GetEventAtOk() (*string, bool)`

GetEventAtOk returns a tuple with the EventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAt

`func (o *LifeCycle) SetEventAt(v string)`

SetEventAt sets EventAt field to given value.


### GetNote

`func (o *LifeCycle) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LifeCycle) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LifeCycle) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *LifeCycle) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *LifeCycle) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *LifeCycle) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *LifeCycle) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


