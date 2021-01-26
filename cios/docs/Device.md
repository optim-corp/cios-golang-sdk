# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ResourceOwnerId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**IdNumber** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayInfo** | Pointer to [**DisplayInfo**](DisplayInfo.md) |  | [optional] 
**ClientList** | Pointer to [**[]Client**](Client.md) |  | [optional] 
**IsManaged** | Pointer to **string** |  | [optional] 
**RsaPublickey** | Pointer to **string** |  | [optional] 
**CustomInventory** | Pointer to **[]interface{}** |  | [optional] 
**CreatedAt** | **string** | 最終更新時間 | 
**UpdatedAt** | **string** | 作成時間 | 
**ActivatedAt** | Pointer to **string** | アクティベート時間。is_managedがfull_managedの場合のみ表示 | [optional] 
**MonitoringStatus** | Pointer to **string** | ok:正常, error:エラー, offline:オフライン、 Null:監視状態なし | [optional] 

## Methods

### NewDevice

`func NewDevice(id string, resourceOwnerId string, createdAt string, updatedAt string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetResourceOwnerId

`func (o *Device) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Device) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Device) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdNumber

`func (o *Device) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *Device) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *Device) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *Device) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### GetDescription

`func (o *Device) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Device) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Device) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Device) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *Device) GetDisplayInfo() DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *Device) GetDisplayInfoOk() (*DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *Device) SetDisplayInfo(v DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *Device) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetClientList

`func (o *Device) GetClientList() []Client`

GetClientList returns the ClientList field if non-nil, zero value otherwise.

### GetClientListOk

`func (o *Device) GetClientListOk() (*[]Client, bool)`

GetClientListOk returns a tuple with the ClientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientList

`func (o *Device) SetClientList(v []Client)`

SetClientList sets ClientList field to given value.

### HasClientList

`func (o *Device) HasClientList() bool`

HasClientList returns a boolean if a field has been set.

### GetIsManaged

`func (o *Device) GetIsManaged() string`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *Device) GetIsManagedOk() (*string, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *Device) SetIsManaged(v string)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *Device) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetRsaPublickey

`func (o *Device) GetRsaPublickey() string`

GetRsaPublickey returns the RsaPublickey field if non-nil, zero value otherwise.

### GetRsaPublickeyOk

`func (o *Device) GetRsaPublickeyOk() (*string, bool)`

GetRsaPublickeyOk returns a tuple with the RsaPublickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPublickey

`func (o *Device) SetRsaPublickey(v string)`

SetRsaPublickey sets RsaPublickey field to given value.

### HasRsaPublickey

`func (o *Device) HasRsaPublickey() bool`

HasRsaPublickey returns a boolean if a field has been set.

### GetCustomInventory

`func (o *Device) GetCustomInventory() []interface{}`

GetCustomInventory returns the CustomInventory field if non-nil, zero value otherwise.

### GetCustomInventoryOk

`func (o *Device) GetCustomInventoryOk() (*[]interface{}, bool)`

GetCustomInventoryOk returns a tuple with the CustomInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInventory

`func (o *Device) SetCustomInventory(v []interface{})`

SetCustomInventory sets CustomInventory field to given value.

### HasCustomInventory

`func (o *Device) HasCustomInventory() bool`

HasCustomInventory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Device) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Device) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Device) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Device) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Device) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Device) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetActivatedAt

`func (o *Device) GetActivatedAt() string`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *Device) GetActivatedAtOk() (*string, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *Device) SetActivatedAt(v string)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *Device) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetMonitoringStatus

`func (o *Device) GetMonitoringStatus() string`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *Device) GetMonitoringStatusOk() (*string, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *Device) SetMonitoringStatus(v string)`

SetMonitoringStatus sets MonitoringStatus field to given value.

### HasMonitoringStatus

`func (o *Device) HasMonitoringStatus() bool`

HasMonitoringStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


