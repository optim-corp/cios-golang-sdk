# DeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOwnerId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**IdNumber** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**ClientIdList** | Pointer to **[]string** | クライアントIDリスト。デバイスコンテキストのクライアントが必要な場合(is_managedをmanagedかfull_managedに設定する場合)、対応するエージェント等、デバイス上で動作させるクライアントのIDと公開鍵を指定する。複数ある場合、すべて登録すること。 | [optional] 
**IsManaged** | Pointer to **string** | 該当デバイスの管理方法についての区分。  管理外：unmanaged ・・・該当デバイス情報に情報を持たせるのみの場合、こちらを選択。 一部管理：managed ・・・該当デバイス情報に情報を持たせ、かつ当サービス以外システム（アプリ）にて管理される場合、こちらを選択。 完全管理：full_managed（次期対応検討内容） ・・・該当デバイス情報に情報を持たせ、かつ当サービス以外システム（アプリ）およびデバイスエージェント（キッティング、ファームウェアアップデート等行う機構）にて管理される場合、こちらを選択。 | [optional] 
**RsaPublickey** | Pointer to **string** | RSAの公開鍵をPEM形式で設定する。クライアントIDを指定する場合必須。 | [optional] 
**CustomInventory** | Pointer to **[]interface{}** | 付帯情報リスト | [optional] 

## Methods

### NewDeviceInfo

`func NewDeviceInfo(resourceOwnerId string, ) *DeviceInfo`

NewDeviceInfo instantiates a new DeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoWithDefaults

`func NewDeviceInfoWithDefaults() *DeviceInfo`

NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOwnerId

`func (o *DeviceInfo) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *DeviceInfo) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *DeviceInfo) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetName

`func (o *DeviceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdNumber

`func (o *DeviceInfo) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *DeviceInfo) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *DeviceInfo) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *DeviceInfo) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *DeviceInfo) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *DeviceInfo) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *DeviceInfo) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *DeviceInfo) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetClientIdList

`func (o *DeviceInfo) GetClientIdList() []string`

GetClientIdList returns the ClientIdList field if non-nil, zero value otherwise.

### GetClientIdListOk

`func (o *DeviceInfo) GetClientIdListOk() (*[]string, bool)`

GetClientIdListOk returns a tuple with the ClientIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdList

`func (o *DeviceInfo) SetClientIdList(v []string)`

SetClientIdList sets ClientIdList field to given value.

### HasClientIdList

`func (o *DeviceInfo) HasClientIdList() bool`

HasClientIdList returns a boolean if a field has been set.

### GetIsManaged

`func (o *DeviceInfo) GetIsManaged() string`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *DeviceInfo) GetIsManagedOk() (*string, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *DeviceInfo) SetIsManaged(v string)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *DeviceInfo) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetRsaPublickey

`func (o *DeviceInfo) GetRsaPublickey() string`

GetRsaPublickey returns the RsaPublickey field if non-nil, zero value otherwise.

### GetRsaPublickeyOk

`func (o *DeviceInfo) GetRsaPublickeyOk() (*string, bool)`

GetRsaPublickeyOk returns a tuple with the RsaPublickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPublickey

`func (o *DeviceInfo) SetRsaPublickey(v string)`

SetRsaPublickey sets RsaPublickey field to given value.

### HasRsaPublickey

`func (o *DeviceInfo) HasRsaPublickey() bool`

HasRsaPublickey returns a boolean if a field has been set.

### GetCustomInventory

`func (o *DeviceInfo) GetCustomInventory() []interface{}`

GetCustomInventory returns the CustomInventory field if non-nil, zero value otherwise.

### GetCustomInventoryOk

`func (o *DeviceInfo) GetCustomInventoryOk() (*[]interface{}, bool)`

GetCustomInventoryOk returns a tuple with the CustomInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInventory

`func (o *DeviceInfo) SetCustomInventory(v []interface{})`

SetCustomInventory sets CustomInventory field to given value.

### HasCustomInventory

`func (o *DeviceInfo) HasCustomInventory() bool`

HasCustomInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


