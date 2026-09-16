# ListNodePoolsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePools** | [**[]NodePool**](NodePool.md) | 클러스터 내 노드 풀 목록 | 

## Methods

### NewListNodePoolsResponse

`func NewListNodePoolsResponse(nodePools []NodePool, ) *ListNodePoolsResponse`

NewListNodePoolsResponse instantiates a new ListNodePoolsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNodePoolsResponseWithDefaults

`func NewListNodePoolsResponseWithDefaults() *ListNodePoolsResponse`

NewListNodePoolsResponseWithDefaults instantiates a new ListNodePoolsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePools

`func (o *ListNodePoolsResponse) GetNodePools() []NodePool`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *ListNodePoolsResponse) GetNodePoolsOk() (*[]NodePool, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *ListNodePoolsResponse) SetNodePools(v []NodePool)`

SetNodePools sets NodePools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


