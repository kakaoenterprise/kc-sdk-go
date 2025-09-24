# GetK8sClusterNodePoolsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePools** | [**[]KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel**](KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel.md) | 클러스터 내 노드 풀 목록 | 

## Methods

### NewGetK8sClusterNodePoolsResponseModel

`func NewGetK8sClusterNodePoolsResponseModel(nodePools []KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel, ) *GetK8sClusterNodePoolsResponseModel`

NewGetK8sClusterNodePoolsResponseModel instantiates a new GetK8sClusterNodePoolsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sClusterNodePoolsResponseModelWithDefaults

`func NewGetK8sClusterNodePoolsResponseModelWithDefaults() *GetK8sClusterNodePoolsResponseModel`

NewGetK8sClusterNodePoolsResponseModelWithDefaults instantiates a new GetK8sClusterNodePoolsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePools

`func (o *GetK8sClusterNodePoolsResponseModel) GetNodePools() []KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *GetK8sClusterNodePoolsResponseModel) GetNodePoolsOk() (*[]KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *GetK8sClusterNodePoolsResponseModel) SetNodePools(v []KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel)`

SetNodePools sets NodePools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


