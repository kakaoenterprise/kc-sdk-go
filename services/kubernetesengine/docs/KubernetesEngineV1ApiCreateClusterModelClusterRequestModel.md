# KubernetesEngineV1ApiCreateClusterModelClusterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 클러스터 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Version** | **string** | 클러스터 Kubernetes 버전 | 
**VpcInfo** | [**VpcInfoRequestModel**](VpcInfoRequestModel.md) | 클러스터가 속한 VPC 정보 | 
**IsAllocateFip** | **bool** | 퍼블릭 IP 할당 여부&lt;br/&gt; - &#x60;true&#x60;: 퍼블릭 IP를 할당&lt;br/&gt; - &#x60;false&#x60;: 퍼블릭 IP를 할당하지 않음 | 
**Network** | [**ClusterNetworkRequestModel**](ClusterNetworkRequestModel.md) | 클러스터 네트워크 구성 정보 | 

## Methods

### NewKubernetesEngineV1ApiCreateClusterModelClusterRequestModel

`func NewKubernetesEngineV1ApiCreateClusterModelClusterRequestModel(name string, version string, vpcInfo VpcInfoRequestModel, isAllocateFip bool, network ClusterNetworkRequestModel, ) *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel`

NewKubernetesEngineV1ApiCreateClusterModelClusterRequestModel instantiates a new KubernetesEngineV1ApiCreateClusterModelClusterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiCreateClusterModelClusterRequestModelWithDefaults

`func NewKubernetesEngineV1ApiCreateClusterModelClusterRequestModelWithDefaults() *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel`

NewKubernetesEngineV1ApiCreateClusterModelClusterRequestModelWithDefaults instantiates a new KubernetesEngineV1ApiCreateClusterModelClusterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVpcInfo

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetVpcInfo() VpcInfoRequestModel`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetVpcInfoOk() (*VpcInfoRequestModel, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) SetVpcInfo(v VpcInfoRequestModel)`

SetVpcInfo sets VpcInfo field to given value.


### GetIsAllocateFip

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetIsAllocateFip() bool`

GetIsAllocateFip returns the IsAllocateFip field if non-nil, zero value otherwise.

### GetIsAllocateFipOk

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetIsAllocateFipOk() (*bool, bool)`

GetIsAllocateFipOk returns a tuple with the IsAllocateFip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllocateFip

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) SetIsAllocateFip(v bool)`

SetIsAllocateFip sets IsAllocateFip field to given value.


### GetNetwork

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetNetwork() ClusterNetworkRequestModel`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) GetNetworkOk() (*ClusterNetworkRequestModel, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *KubernetesEngineV1ApiCreateClusterModelClusterRequestModel) SetNetwork(v ClusterNetworkRequestModel)`

SetNetwork sets Network field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


