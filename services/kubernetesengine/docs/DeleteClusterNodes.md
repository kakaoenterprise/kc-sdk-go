# DeleteClusterNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRemove** | **bool** | 노드 삭제/재생성 여부&lt;br/&gt;- &#x60;true&#x60;:  해당 노드를 클러스터에서 삭제 &lt;br/&gt;- &#x60;false&#x60;: 해당 노드를 재생성 (노드 재생성 시, 해당 노드를 drain 후 새로운 노드를 생성하며 기존 노드는 삭제됨) | 
**NodeNames** | **[]string** | 제거할 노드 이름 목록 | 

## Methods

### NewDeleteClusterNodes

`func NewDeleteClusterNodes(isRemove bool, nodeNames []string, ) *DeleteClusterNodes`

NewDeleteClusterNodes instantiates a new DeleteClusterNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteClusterNodesWithDefaults

`func NewDeleteClusterNodesWithDefaults() *DeleteClusterNodes`

NewDeleteClusterNodesWithDefaults instantiates a new DeleteClusterNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRemove

`func (o *DeleteClusterNodes) GetIsRemove() bool`

GetIsRemove returns the IsRemove field if non-nil, zero value otherwise.

### GetIsRemoveOk

`func (o *DeleteClusterNodes) GetIsRemoveOk() (*bool, bool)`

GetIsRemoveOk returns a tuple with the IsRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemove

`func (o *DeleteClusterNodes) SetIsRemove(v bool)`

SetIsRemove sets IsRemove field to given value.


### GetNodeNames

`func (o *DeleteClusterNodes) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *DeleteClusterNodes) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *DeleteClusterNodes) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


