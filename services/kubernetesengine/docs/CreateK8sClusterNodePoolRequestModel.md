# CreateK8sClusterNodePoolRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePool** | [**KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel**](KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel.md) | 대상 노드 풀 | 

## Methods

### NewCreateK8sClusterNodePoolRequestModel

`func NewCreateK8sClusterNodePoolRequestModel(nodePool KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel, ) *CreateK8sClusterNodePoolRequestModel`

NewCreateK8sClusterNodePoolRequestModel instantiates a new CreateK8sClusterNodePoolRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateK8sClusterNodePoolRequestModelWithDefaults

`func NewCreateK8sClusterNodePoolRequestModelWithDefaults() *CreateK8sClusterNodePoolRequestModel`

NewCreateK8sClusterNodePoolRequestModelWithDefaults instantiates a new CreateK8sClusterNodePoolRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePool

`func (o *CreateK8sClusterNodePoolRequestModel) GetNodePool() KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel`

GetNodePool returns the NodePool field if non-nil, zero value otherwise.

### GetNodePoolOk

`func (o *CreateK8sClusterNodePoolRequestModel) GetNodePoolOk() (*KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel, bool)`

GetNodePoolOk returns a tuple with the NodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePool

`func (o *CreateK8sClusterNodePoolRequestModel) SetNodePool(v KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel)`

SetNodePool sets NodePool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


