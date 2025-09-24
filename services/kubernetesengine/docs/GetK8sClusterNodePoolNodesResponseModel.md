# GetK8sClusterNodePoolNodesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel**](KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel.md) | 노드 목록 | 

## Methods

### NewGetK8sClusterNodePoolNodesResponseModel

`func NewGetK8sClusterNodePoolNodesResponseModel(nodes []KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel, ) *GetK8sClusterNodePoolNodesResponseModel`

NewGetK8sClusterNodePoolNodesResponseModel instantiates a new GetK8sClusterNodePoolNodesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sClusterNodePoolNodesResponseModelWithDefaults

`func NewGetK8sClusterNodePoolNodesResponseModelWithDefaults() *GetK8sClusterNodePoolNodesResponseModel`

NewGetK8sClusterNodePoolNodesResponseModelWithDefaults instantiates a new GetK8sClusterNodePoolNodesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *GetK8sClusterNodePoolNodesResponseModel) GetNodes() []KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetK8sClusterNodePoolNodesResponseModel) GetNodesOk() (*[]KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetK8sClusterNodePoolNodesResponseModel) SetNodes(v []KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


