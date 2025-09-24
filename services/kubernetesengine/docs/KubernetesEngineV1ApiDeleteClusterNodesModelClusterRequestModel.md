# KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRemove** | **bool** | 노드 삭제/재생성 여부&lt;br/&gt;- &#x60;true&#x60;:  해당 노드를 클러스터에서 삭제 &lt;br/&gt;- &#x60;false&#x60;: 해당 노드를 재생성 (노드 재생성 시, 해당 노드를 drain 후 새로운 노드를 생성하며 기존 노드는 삭제됨) | 
**NodeNames** | **[]string** | 제거할 노드 이름 목록 | 

## Methods

### NewKubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel

`func NewKubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel(isRemove bool, nodeNames []string, ) *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel`

NewKubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel instantiates a new KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModelWithDefaults

`func NewKubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModelWithDefaults() *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel`

NewKubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModelWithDefaults instantiates a new KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRemove

`func (o *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel) GetIsRemove() bool`

GetIsRemove returns the IsRemove field if non-nil, zero value otherwise.

### GetIsRemoveOk

`func (o *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel) GetIsRemoveOk() (*bool, bool)`

GetIsRemoveOk returns a tuple with the IsRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemove

`func (o *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel) SetIsRemove(v bool)`

SetIsRemove sets IsRemove field to given value.


### GetNodeNames

`func (o *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


