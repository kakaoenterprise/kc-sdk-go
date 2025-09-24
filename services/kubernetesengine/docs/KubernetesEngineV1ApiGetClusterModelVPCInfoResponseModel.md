# KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | [**[]KubernetesEngineV1ApiGetClusterModelSubnetResponseModel**](KubernetesEngineV1ApiGetClusterModelSubnetResponseModel.md) | 서브넷 정보 목록 | 

## Methods

### NewKubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel

`func NewKubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel(id string, subnets []KubernetesEngineV1ApiGetClusterModelSubnetResponseModel, ) *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel`

NewKubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel instantiates a new KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiGetClusterModelVPCInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiGetClusterModelVPCInfoResponseModelWithDefaults() *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel`

NewKubernetesEngineV1ApiGetClusterModelVPCInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel) GetSubnets() []KubernetesEngineV1ApiGetClusterModelSubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel) GetSubnetsOk() (*[]KubernetesEngineV1ApiGetClusterModelSubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel) SetSubnets(v []KubernetesEngineV1ApiGetClusterModelSubnetResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


