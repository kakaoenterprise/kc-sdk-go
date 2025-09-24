# GetK8sClusterNodesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel**](KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel.md) | 클러스터 내 노드 목록 | 

## Methods

### NewGetK8sClusterNodesResponseModel

`func NewGetK8sClusterNodesResponseModel(nodes []KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel, ) *GetK8sClusterNodesResponseModel`

NewGetK8sClusterNodesResponseModel instantiates a new GetK8sClusterNodesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sClusterNodesResponseModelWithDefaults

`func NewGetK8sClusterNodesResponseModelWithDefaults() *GetK8sClusterNodesResponseModel`

NewGetK8sClusterNodesResponseModelWithDefaults instantiates a new GetK8sClusterNodesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *GetK8sClusterNodesResponseModel) GetNodes() []KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetK8sClusterNodesResponseModel) GetNodesOk() (*[]KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetK8sClusterNodesResponseModel) SetNodes(v []KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


