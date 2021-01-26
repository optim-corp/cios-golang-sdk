# ConstitutionComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ComponentTypeEnum**](ComponentTypeEnum.md) |  | [optional] 
**Attribute** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayInfo** | Pointer to [**[]DisplayInfo**](DisplayInfo.md) |  | [optional] 
**Components** | Pointer to [**[]ConstitutionComponent**](ConstitutionComponent.md) |  | [optional] 

## Methods

### NewConstitutionComponent

`func NewConstitutionComponent() *ConstitutionComponent`

NewConstitutionComponent instantiates a new ConstitutionComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstitutionComponentWithDefaults

`func NewConstitutionComponentWithDefaults() *ConstitutionComponent`

NewConstitutionComponentWithDefaults instantiates a new ConstitutionComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConstitutionComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConstitutionComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConstitutionComponent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConstitutionComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *ConstitutionComponent) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ConstitutionComponent) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ConstitutionComponent) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ConstitutionComponent) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAddress

`func (o *ConstitutionComponent) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConstitutionComponent) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConstitutionComponent) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConstitutionComponent) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetType

`func (o *ConstitutionComponent) GetType() ComponentTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstitutionComponent) GetTypeOk() (*ComponentTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstitutionComponent) SetType(v ComponentTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *ConstitutionComponent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttribute

`func (o *ConstitutionComponent) GetAttribute() map[string]interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ConstitutionComponent) GetAttributeOk() (*map[string]interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ConstitutionComponent) SetAttribute(v map[string]interface{})`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *ConstitutionComponent) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *ConstitutionComponent) GetDisplayInfo() []DisplayInfo`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *ConstitutionComponent) GetDisplayInfoOk() (*[]DisplayInfo, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *ConstitutionComponent) SetDisplayInfo(v []DisplayInfo)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *ConstitutionComponent) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetComponents

`func (o *ConstitutionComponent) GetComponents() []ConstitutionComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ConstitutionComponent) GetComponentsOk() (*[]ConstitutionComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ConstitutionComponent) SetComponents(v []ConstitutionComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ConstitutionComponent) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


