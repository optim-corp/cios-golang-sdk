# SubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionItem

`func NewSubscriptionItem() *SubscriptionItem`

NewSubscriptionItem instantiates a new SubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemWithDefaults

`func NewSubscriptionItemWithDefaults() *SubscriptionItem`

NewSubscriptionItemWithDefaults instantiates a new SubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *SubscriptionItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SubscriptionItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SubscriptionItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SubscriptionItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionItem) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionItem) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionItem) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionItem) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


