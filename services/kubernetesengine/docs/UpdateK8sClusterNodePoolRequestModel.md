# UpdateK8sClusterNodePoolRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePool** | [**KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel**](KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel.md) | 대상 노드 풀 | 

## Methods

### NewUpdateK8sClusterNodePoolRequestModel

`func NewUpdateK8sClusterNodePoolRequestModel(nodePool KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel, ) *UpdateK8sClusterNodePoolRequestModel`

NewUpdateK8sClusterNodePoolRequestModel instantiates a new UpdateK8sClusterNodePoolRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateK8sClusterNodePoolRequestModelWithDefaults

`func NewUpdateK8sClusterNodePoolRequestModelWithDefaults() *UpdateK8sClusterNodePoolRequestModel`

NewUpdateK8sClusterNodePoolRequestModelWithDefaults instantiates a new UpdateK8sClusterNodePoolRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePool

`func (o *UpdateK8sClusterNodePoolRequestModel) GetNodePool() KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel`

GetNodePool returns the NodePool field if non-nil, zero value otherwise.

### GetNodePoolOk

`func (o *UpdateK8sClusterNodePoolRequestModel) GetNodePoolOk() (*KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel, bool)`

GetNodePoolOk returns a tuple with the NodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePool

`func (o *UpdateK8sClusterNodePoolRequestModel) SetNodePool(v KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel)`

SetNodePool sets NodePool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


