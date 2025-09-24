# KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | [**[]KubernetesEngineV1ApiGetClusterNodeModelSubnetResponseModel**](KubernetesEngineV1ApiGetClusterNodeModelSubnetResponseModel.md) | 서브넷 ID 목록 | 

## Methods

### NewKubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel

`func NewKubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel(id string, subnets []KubernetesEngineV1ApiGetClusterNodeModelSubnetResponseModel, ) *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel instantiates a new KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModelWithDefaults() *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel) GetSubnets() []KubernetesEngineV1ApiGetClusterNodeModelSubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel) GetSubnetsOk() (*[]KubernetesEngineV1ApiGetClusterNodeModelSubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *KubernetesEngineV1ApiGetClusterNodeModelVpcInfoResponseModel) SetSubnets(v []KubernetesEngineV1ApiGetClusterNodeModelSubnetResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


