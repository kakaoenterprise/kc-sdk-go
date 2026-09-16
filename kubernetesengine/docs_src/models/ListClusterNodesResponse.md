# ListClusterNodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]Node**](Node.md) | 클러스터 내 노드 목록 | 

## Methods

### NewListClusterNodesResponse

`func NewListClusterNodesResponse(nodes []Node, ) *ListClusterNodesResponse`

NewListClusterNodesResponse instantiates a new ListClusterNodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterNodesResponseWithDefaults

`func NewListClusterNodesResponseWithDefaults() *ListClusterNodesResponse`

NewListClusterNodesResponseWithDefaults instantiates a new ListClusterNodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ListClusterNodesResponse) GetNodes() []Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ListClusterNodesResponse) GetNodesOk() (*[]Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ListClusterNodesResponse) SetNodes(v []Node)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


