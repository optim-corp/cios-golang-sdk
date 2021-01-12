# UpdatedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**CreatedByUser**](CreatedBy_user.md) |  | [optional] 
**Client** | Pointer to [**CreatedByClient**](CreatedBy_client.md) |  | [optional] 

## Methods

### NewUpdatedBy

`func NewUpdatedBy() *UpdatedBy`

NewUpdatedBy instantiates a new UpdatedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedByWithDefaults

`func NewUpdatedByWithDefaults() *UpdatedBy`

NewUpdatedByWithDefaults instantiates a new UpdatedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UpdatedBy) GetUser() CreatedByUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdatedBy) GetUserOk() (*CreatedByUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdatedBy) SetUser(v CreatedByUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdatedBy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetClient

`func (o *UpdatedBy) GetClient() CreatedByClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *UpdatedBy) GetClientOk() (*CreatedByClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *UpdatedBy) SetClient(v CreatedByClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *UpdatedBy) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


