# KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | [**[]KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel**](KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel.md) | 서브넷 정보 | 

## Methods

### NewKubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel

`func NewKubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel(id string, subnets []KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel, ) *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel instantiates a new KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModelWithDefaults() *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel`

NewKubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel) GetSubnets() []KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel) GetSubnetsOk() (*[]KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel) SetSubnets(v []KubernetesEngineV1ApiListClusterNodesModelSubnetResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


