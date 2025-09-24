# KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | [**[]KubernetesEngineV1ApiListClustersModelSubnetResponseModel**](KubernetesEngineV1ApiListClustersModelSubnetResponseModel.md) | 서브넷 정보 | 

## Methods

### NewKubernetesEngineV1ApiListClustersModelVPCInfoResponseModel

`func NewKubernetesEngineV1ApiListClustersModelVPCInfoResponseModel(id string, subnets []KubernetesEngineV1ApiListClustersModelSubnetResponseModel, ) *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel`

NewKubernetesEngineV1ApiListClustersModelVPCInfoResponseModel instantiates a new KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListClustersModelVPCInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListClustersModelVPCInfoResponseModelWithDefaults() *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel`

NewKubernetesEngineV1ApiListClustersModelVPCInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel) GetSubnets() []KubernetesEngineV1ApiListClustersModelSubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel) GetSubnetsOk() (*[]KubernetesEngineV1ApiListClustersModelSubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *KubernetesEngineV1ApiListClustersModelVPCInfoResponseModel) SetSubnets(v []KubernetesEngineV1ApiListClustersModelSubnetResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


