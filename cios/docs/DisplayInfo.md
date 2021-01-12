# DisplayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Language** | **string** |  | 
**Script** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**IsDefault** | **bool** |  | 

## Methods

### NewDisplayInfo

`func NewDisplayInfo(name string, language string, isDefault bool, ) *DisplayInfo`

NewDisplayInfo instantiates a new DisplayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayInfoWithDefaults

`func NewDisplayInfoWithDefaults() *DisplayInfo`

NewDisplayInfoWithDefaults instantiates a new DisplayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DisplayInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DisplayInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DisplayInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DisplayInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DisplayInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DisplayInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DisplayInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLanguage

`func (o *DisplayInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DisplayInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DisplayInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetScript

`func (o *DisplayInfo) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *DisplayInfo) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *DisplayInfo) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *DisplayInfo) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetRegion

`func (o *DisplayInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DisplayInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DisplayInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DisplayInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIsDefault

`func (o *DisplayInfo) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DisplayInfo) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DisplayInfo) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


