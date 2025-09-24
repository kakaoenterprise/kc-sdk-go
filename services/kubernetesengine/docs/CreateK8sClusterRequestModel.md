# CreateK8sClusterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**KubernetesEngineV1ApiCreateClusterModelClusterRequestModel**](KubernetesEngineV1ApiCreateClusterModelClusterRequestModel.md) | 대상 클러스터 정보 | 

## Methods

### NewCreateK8sClusterRequestModel

`func NewCreateK8sClusterRequestModel(cluster KubernetesEngineV1ApiCreateClusterModelClusterRequestModel, ) *CreateK8sClusterRequestModel`

NewCreateK8sClusterRequestModel instantiates a new CreateK8sClusterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateK8sClusterRequestModelWithDefaults

`func NewCreateK8sClusterRequestModelWithDefaults() *CreateK8sClusterRequestModel`

NewCreateK8sClusterRequestModelWithDefaults instantiates a new CreateK8sClusterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *CreateK8sClusterRequestModel) GetCluster() KubernetesEngineV1ApiCreateClusterModelClusterRequestModel`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CreateK8sClusterRequestModel) GetClusterOk() (*KubernetesEngineV1ApiCreateClusterModelClusterRequestModel, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CreateK8sClusterRequestModel) SetCluster(v KubernetesEngineV1ApiCreateClusterModelClusterRequestModel)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


