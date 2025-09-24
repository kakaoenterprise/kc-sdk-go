# KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | [**[]KubernetesEngineV1ApiListNodePoolsModelSubnetResponseModel**](KubernetesEngineV1ApiListNodePoolsModelSubnetResponseModel.md) | 서브넷 정보 목록 | 

## Methods

### NewKubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel

`func NewKubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel(id string, subnets []KubernetesEngineV1ApiListNodePoolsModelSubnetResponseModel, ) *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel instantiates a new KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModelWithDefaults() *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel) GetSubnets() []KubernetesEngineV1ApiListNodePoolsModelSubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel) GetSubnetsOk() (*[]KubernetesEngineV1ApiListNodePoolsModelSubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel) SetSubnets(v []KubernetesEngineV1ApiListNodePoolsModelSubnetResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


