# GetK8sClusterNodePoolResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePool** | [**KubernetesEngineV1ApiGetNodePoolModelNodePoolResponseModel**](KubernetesEngineV1ApiGetNodePoolModelNodePoolResponseModel.md) | 노드 풀 정보 | 

## Methods

### NewGetK8sClusterNodePoolResponseModel

`func NewGetK8sClusterNodePoolResponseModel(nodePool KubernetesEngineV1ApiGetNodePoolModelNodePoolResponseModel, ) *GetK8sClusterNodePoolResponseModel`

NewGetK8sClusterNodePoolResponseModel instantiates a new GetK8sClusterNodePoolResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sClusterNodePoolResponseModelWithDefaults

`func NewGetK8sClusterNodePoolResponseModelWithDefaults() *GetK8sClusterNodePoolResponseModel`

NewGetK8sClusterNodePoolResponseModelWithDefaults instantiates a new GetK8sClusterNodePoolResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePool

`func (o *GetK8sClusterNodePoolResponseModel) GetNodePool() KubernetesEngineV1ApiGetNodePoolModelNodePoolResponseModel`

GetNodePool returns the NodePool field if non-nil, zero value otherwise.

### GetNodePoolOk

`func (o *GetK8sClusterNodePoolResponseModel) GetNodePoolOk() (*KubernetesEngineV1ApiGetNodePoolModelNodePoolResponseModel, bool)`

GetNodePoolOk returns a tuple with the NodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePool

`func (o *GetK8sClusterNodePoolResponseModel) SetNodePool(v KubernetesEngineV1ApiGetNodePoolModelNodePoolResponseModel)`

SetNodePool sets NodePool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


