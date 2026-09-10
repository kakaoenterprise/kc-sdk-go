# CreateNodePoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePool** | [**CreateNodePool**](CreateNodePool.md) | 대상 노드 풀 | 

## Methods

### NewCreateNodePoolRequest

`func NewCreateNodePoolRequest(nodePool CreateNodePool, ) *CreateNodePoolRequest`

NewCreateNodePoolRequest instantiates a new CreateNodePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNodePoolRequestWithDefaults

`func NewCreateNodePoolRequestWithDefaults() *CreateNodePoolRequest`

NewCreateNodePoolRequestWithDefaults instantiates a new CreateNodePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePool

`func (o *CreateNodePoolRequest) GetNodePool() CreateNodePool`

GetNodePool returns the NodePool field if non-nil, zero value otherwise.

### GetNodePoolOk

`func (o *CreateNodePoolRequest) GetNodePoolOk() (*CreateNodePool, bool)`

GetNodePoolOk returns a tuple with the NodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePool

`func (o *CreateNodePoolRequest) SetNodePool(v CreateNodePool)`

SetNodePool sets NodePool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


