# KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | [**[]KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel**](KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel.md) | 서브넷 ID 목록 | 

## Methods

### NewKubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel

`func NewKubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel(id string, subnets []KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel, ) *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel instantiates a new KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModelWithDefaults() *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel) GetSubnets() []KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel) GetSubnetsOk() (*[]KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *KubernetesEngineV1ApiGetNodePoolModelVpcInfoResponseModel) SetSubnets(v []KubernetesEngineV1ApiGetNodePoolModelSubnetResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


