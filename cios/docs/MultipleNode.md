# MultipleNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** |  | 
**Nodes** | [**[]Node**](Node.md) |  | 

## Methods

### NewMultipleNode

`func NewMultipleNode(total int64, nodes []Node, ) *MultipleNode`

NewMultipleNode instantiates a new MultipleNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleNodeWithDefaults

`func NewMultipleNodeWithDefaults() *MultipleNode`

NewMultipleNodeWithDefaults instantiates a new MultipleNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MultipleNode) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MultipleNode) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MultipleNode) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetNodes

`func (o *MultipleNode) GetNodes() []Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *MultipleNode) GetNodesOk() (*[]Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *MultipleNode) SetNodes(v []Node)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


