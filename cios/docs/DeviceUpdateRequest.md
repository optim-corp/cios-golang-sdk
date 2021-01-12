# DeviceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IdNumber** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayInfo** | Pointer to [**DisplayInfo**](DisplayInfo.md) |  | [optional] 
**ClientIdList** | Pointer to **[]string** |  | [optional] 
**IsManaged** | Pointer to **string** | 該当デバイスの管理方法についての区分。8  管理外：unmanaged ・・・該当デバイス情報に情報を持たせるのみの場合、こちらを選択。 一部管理：managed ・・・該当デバイス情報に情報を持たせ、かつ当サービス以外システム（アプリ）にて管理される場合、こちらを選択。 完全管理：full_managed ・・・該当デバイス情報に情報を持たせ、かつ当サービス以外システム（アプリ）およびデバイスエージェント（キッティング、ファームウェアアップデート等行う機構）にて管理される場合、こちらを選択。 | [optional] 
**RsaPublickey** | Pointer to **string** |  | [optional] 
**CustomInventory** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewDeviceUpdateRequest

`func NewDeviceUpdateRequest() *DeviceUpdateRequest`

NewDeviceUpdateRequest instantiates a new DeviceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpdateRequestWithDefaults

`func NewDeviceUpdateRequestWithDefaults() *DeviceUpdateRequest`

NewDeviceUpdateRequestWithDefaults instantiates a new DeviceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdNumber

`func (o *DeviceUpdateRequest) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *DeviceUpdateRequest) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *DeviceUpdateRequest) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *DeviceUpdateRequest) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *DeviceUpdateRequest) GetDisplayInfo() DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *DeviceUpdateRequest) GetDisplayInfoOk() (*DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *DeviceUpdateRequest) SetDisplayInfo(v DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *DeviceUpdateRequest) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetClientIdList

`func (o *DeviceUpdateRequest) GetClientIdList() []string`

GetClientIdList returns the ClientIdList field if non-nil, zero value otherwise.

### GetClientIdListOk

`func (o *DeviceUpdateRequest) GetClientIdListOk() (*[]string, bool)`

GetClientIdListOk returns a tuple with the ClientIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdList

`func (o *DeviceUpdateRequest) SetClientIdList(v []string)`

SetClientIdList sets ClientIdList field to given value.

### HasClientIdList

`func (o *DeviceUpdateRequest) HasClientIdList() bool`

HasClientIdList returns a boolean if a field has been set.

### GetIsManaged

`func (o *DeviceUpdateRequest) GetIsManaged() string`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *DeviceUpdateRequest) GetIsManagedOk() (*string, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *DeviceUpdateRequest) SetIsManaged(v string)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *DeviceUpdateRequest) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetRsaPublickey

`func (o *DeviceUpdateRequest) GetRsaPublickey() string`

GetRsaPublickey returns the RsaPublickey field if non-nil, zero value otherwise.

### GetRsaPublickeyOk

`func (o *DeviceUpdateRequest) GetRsaPublickeyOk() (*string, bool)`

GetRsaPublickeyOk returns a tuple with the RsaPublickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPublickey

`func (o *DeviceUpdateRequest) SetRsaPublickey(v string)`

SetRsaPublickey sets RsaPublickey field to given value.

### HasRsaPublickey

`func (o *DeviceUpdateRequest) HasRsaPublickey() bool`

HasRsaPublickey returns a boolean if a field has been set.

### GetCustomInventory

`func (o *DeviceUpdateRequest) GetCustomInventory() []interface{}`

GetCustomInventory returns the CustomInventory field if non-nil, zero value otherwise.

### GetCustomInventoryOk

`func (o *DeviceUpdateRequest) GetCustomInventoryOk() (*[]interface{}, bool)`

GetCustomInventoryOk returns a tuple with the CustomInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInventory

`func (o *DeviceUpdateRequest) SetCustomInventory(v []interface{})`

SetCustomInventory sets CustomInventory field to given value.

### HasCustomInventory

`func (o *DeviceUpdateRequest) HasCustomInventory() bool`

HasCustomInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


