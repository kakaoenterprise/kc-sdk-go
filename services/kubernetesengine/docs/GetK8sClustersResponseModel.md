# GetK8sClustersResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | [**[]KubernetesEngineV1ApiListClustersModelClusterResponseModel**](KubernetesEngineV1ApiListClustersModelClusterResponseModel.md) | 클러스터 목록 | 

## Methods

### NewGetK8sClustersResponseModel

`func NewGetK8sClustersResponseModel(clusters []KubernetesEngineV1ApiListClustersModelClusterResponseModel, ) *GetK8sClustersResponseModel`

NewGetK8sClustersResponseModel instantiates a new GetK8sClustersResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sClustersResponseModelWithDefaults

`func NewGetK8sClustersResponseModelWithDefaults() *GetK8sClustersResponseModel`

NewGetK8sClustersResponseModelWithDefaults instantiates a new GetK8sClustersResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *GetK8sClustersResponseModel) GetClusters() []KubernetesEngineV1ApiListClustersModelClusterResponseModel`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GetK8sClustersResponseModel) GetClustersOk() (*[]KubernetesEngineV1ApiListClustersModelClusterResponseModel, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GetK8sClustersResponseModel) SetClusters(v []KubernetesEngineV1ApiListClustersModelClusterResponseModel)`

SetClusters sets Clusters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


