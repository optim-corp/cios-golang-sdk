# Inventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | Pointer to **string** |  | [optional] 
**ResourceOwnerId** | Pointer to **string** |  | [optional] 
**StartAt** | Pointer to **string** |  | [optional] 
**CustomInventory** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInventory

`func NewInventory() *Inventory`

NewInventory instantiates a new Inventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryWithDefaults

`func NewInventoryWithDefaults() *Inventory`

NewInventoryWithDefaults instantiates a new Inventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *Inventory) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Inventory) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Inventory) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Inventory) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *Inventory) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Inventory) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Inventory) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.

### HasResourceOwnerId

`func (o *Inventory) HasResourceOwnerId() bool`

HasResourceOwnerId returns a boolean if a field has been set.

### GetStartAt

`func (o *Inventory) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *Inventory) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *Inventory) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *Inventory) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetCustomInventory

`func (o *Inventory) GetCustomInventory() map[string]interface{}`

GetCustomInventory returns the CustomInventory field if non-nil, zero value otherwise.

### GetCustomInventoryOk

`func (o *Inventory) GetCustomInventoryOk() (*map[string]interface{}, bool)`

GetCustomInventoryOk returns a tuple with the CustomInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInventory

`func (o *Inventory) SetCustomInventory(v map[string]interface{})`

SetCustomInventory sets CustomInventory field to given value.

### HasCustomInventory

`func (o *Inventory) HasCustomInventory() bool`

HasCustomInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


