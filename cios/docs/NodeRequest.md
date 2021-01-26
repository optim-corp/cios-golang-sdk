# NodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ParentNodeId** | Pointer to **string** |  | [optional] 

## Methods

### NewNodeRequest

`func NewNodeRequest(name string, ) *NodeRequest`

NewNodeRequest instantiates a new NodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeRequestWithDefaults

`func NewNodeRequestWithDefaults() *NodeRequest`

NewNodeRequestWithDefaults instantiates a new NodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentNodeId

`func (o *NodeRequest) GetParentNodeId() string`

GetParentNodeId returns the ParentNodeId field if non-nil, zero value otherwise.

### GetParentNodeIdOk

`func (o *NodeRequest) GetParentNodeIdOk() (*string, bool)`

GetParentNodeIdOk returns a tuple with the ParentNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNodeId

`func (o *NodeRequest) SetParentNodeId(v string)`

SetParentNodeId sets ParentNodeId field to given value.

### HasParentNodeId

`func (o *NodeRequest) HasParentNodeId() bool`

HasParentNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


