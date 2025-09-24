# KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | [**[]KubernetesEngineV1ApiListNodePoolNodesModelSubnetResponseModel**](KubernetesEngineV1ApiListNodePoolNodesModelSubnetResponseModel.md) | 서브넷 정보 목록 | 

## Methods

### NewKubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel

`func NewKubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel(id string, subnets []KubernetesEngineV1ApiListNodePoolNodesModelSubnetResponseModel, ) *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel instantiates a new KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModelWithDefaults() *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel) GetSubnets() []KubernetesEngineV1ApiListNodePoolNodesModelSubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel) GetSubnetsOk() (*[]KubernetesEngineV1ApiListNodePoolNodesModelSubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel) SetSubnets(v []KubernetesEngineV1ApiListNodePoolNodesModelSubnetResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


