# MultipleLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Licenses** | [**[]License**](License.md) |  | 

## Methods

### NewMultipleLicense

`func NewMultipleLicense(total int64, licenses []License, ) *MultipleLicense`

NewMultipleLicense instantiates a new MultipleLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleLicenseWithDefaults

`func NewMultipleLicenseWithDefaults() *MultipleLicense`

NewMultipleLicenseWithDefaults instantiates a new MultipleLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleLicense) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleLicense) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleLicense) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetLicenses

`func (o *MultipleLicense) GetLicenses() []License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *MultipleLicense) GetLicensesOk() (*[]License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *MultipleLicense) SetLicenses(v []License)`

SetLicenses sets Licenses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


