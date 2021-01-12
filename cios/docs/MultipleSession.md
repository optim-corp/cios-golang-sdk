# MultipleSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Sessions** | Pointer to [**[]Session**](Session.md) |  | [optional] 

## Methods

### NewMultipleSession

`func NewMultipleSession() *MultipleSession`

NewMultipleSession instantiates a new MultipleSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleSessionWithDefaults

`func NewMultipleSessionWithDefaults() *MultipleSession`

NewMultipleSessionWithDefaults instantiates a new MultipleSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleSession) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleSession) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleSession) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MultipleSession) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSessions

`func (o *MultipleSession) GetSessions() []Session`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *MultipleSession) GetSessionsOk() (*[]Session, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *MultipleSession) SetSessions(v []Session)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *MultipleSession) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


