# UpdateK8sClusterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**KubernetesEngineV1ApiUpdateClusterModelClusterRequestModel**](KubernetesEngineV1ApiUpdateClusterModelClusterRequestModel.md) | 대상 클러스터 이름 | 

## Methods

### NewUpdateK8sClusterRequestModel

`func NewUpdateK8sClusterRequestModel(cluster KubernetesEngineV1ApiUpdateClusterModelClusterRequestModel, ) *UpdateK8sClusterRequestModel`

NewUpdateK8sClusterRequestModel instantiates a new UpdateK8sClusterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateK8sClusterRequestModelWithDefaults

`func NewUpdateK8sClusterRequestModelWithDefaults() *UpdateK8sClusterRequestModel`

NewUpdateK8sClusterRequestModelWithDefaults instantiates a new UpdateK8sClusterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *UpdateK8sClusterRequestModel) GetCluster() KubernetesEngineV1ApiUpdateClusterModelClusterRequestModel`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateK8sClusterRequestModel) GetClusterOk() (*KubernetesEngineV1ApiUpdateClusterModelClusterRequestModel, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateK8sClusterRequestModel) SetCluster(v KubernetesEngineV1ApiUpdateClusterModelClusterRequestModel)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


