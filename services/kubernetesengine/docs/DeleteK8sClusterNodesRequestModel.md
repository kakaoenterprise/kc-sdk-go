# DeleteK8sClusterNodesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel**](KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel.md) | 대상 클러스터 정보 | 

## Methods

### NewDeleteK8sClusterNodesRequestModel

`func NewDeleteK8sClusterNodesRequestModel(cluster KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel, ) *DeleteK8sClusterNodesRequestModel`

NewDeleteK8sClusterNodesRequestModel instantiates a new DeleteK8sClusterNodesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteK8sClusterNodesRequestModelWithDefaults

`func NewDeleteK8sClusterNodesRequestModelWithDefaults() *DeleteK8sClusterNodesRequestModel`

NewDeleteK8sClusterNodesRequestModelWithDefaults instantiates a new DeleteK8sClusterNodesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DeleteK8sClusterNodesRequestModel) GetCluster() KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeleteK8sClusterNodesRequestModel) GetClusterOk() (*KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeleteK8sClusterNodesRequestModel) SetCluster(v KubernetesEngineV1ApiDeleteClusterNodesModelClusterRequestModel)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


