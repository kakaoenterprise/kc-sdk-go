# CreateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 클러스터 이름 | 
**Description** | Pointer to **NullableString** | 클러스터에 대한 설명 | [optional] 
**Version** | **string** | 클러스터 Kubernetes 버전 | 
**VpcInfo** | [**VpcInfoRequest**](VpcInfoRequest.md) | 클러스터가 속한 VPC 정보 | 
**IsAllocateFip** | **bool** | 퍼블릭 IP 할당 여부 - &#x60;true&#x60;: 퍼블릭 IP를 할당 - &#x60;false&#x60;: 퍼블릭 IP를 할당하지 않음 | 
**Network** | [**ClusterNetworkRequest**](ClusterNetworkRequest.md) | 클러스터 네트워크 구성 정보 | 
**ApiServerIngressRules** | Pointer to [**[]ApiServerIngressRuleRequest**](ApiServerIngressRuleRequest.md) | API 서버에 접근을 허용할 CIDR 목록 | [optional] 

## Methods

### NewCreateCluster

`func NewCreateCluster(name string, version string, vpcInfo VpcInfoRequest, isAllocateFip bool, network ClusterNetworkRequest, ) *CreateCluster`

NewCreateCluster instantiates a new CreateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterWithDefaults

`func NewCreateClusterWithDefaults() *CreateCluster`

NewCreateClusterWithDefaults instantiates a new CreateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateCluster) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateCluster) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *CreateCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateCluster) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVpcInfo

`func (o *CreateCluster) GetVpcInfo() VpcInfoRequest`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *CreateCluster) GetVpcInfoOk() (*VpcInfoRequest, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *CreateCluster) SetVpcInfo(v VpcInfoRequest)`

SetVpcInfo sets VpcInfo field to given value.


### GetIsAllocateFip

`func (o *CreateCluster) GetIsAllocateFip() bool`

GetIsAllocateFip returns the IsAllocateFip field if non-nil, zero value otherwise.

### GetIsAllocateFipOk

`func (o *CreateCluster) GetIsAllocateFipOk() (*bool, bool)`

GetIsAllocateFipOk returns a tuple with the IsAllocateFip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllocateFip

`func (o *CreateCluster) SetIsAllocateFip(v bool)`

SetIsAllocateFip sets IsAllocateFip field to given value.


### GetNetwork

`func (o *CreateCluster) GetNetwork() ClusterNetworkRequest`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateCluster) GetNetworkOk() (*ClusterNetworkRequest, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateCluster) SetNetwork(v ClusterNetworkRequest)`

SetNetwork sets Network field to given value.


### GetApiServerIngressRules

`func (o *CreateCluster) GetApiServerIngressRules() []ApiServerIngressRuleRequest`

GetApiServerIngressRules returns the ApiServerIngressRules field if non-nil, zero value otherwise.

### GetApiServerIngressRulesOk

`func (o *CreateCluster) GetApiServerIngressRulesOk() (*[]ApiServerIngressRuleRequest, bool)`

GetApiServerIngressRulesOk returns a tuple with the ApiServerIngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerIngressRules

`func (o *CreateCluster) SetApiServerIngressRules(v []ApiServerIngressRuleRequest)`

SetApiServerIngressRules sets ApiServerIngressRules field to given value.

### HasApiServerIngressRules

`func (o *CreateCluster) HasApiServerIngressRules() bool`

HasApiServerIngressRules returns a boolean if a field has been set.

### SetApiServerIngressRulesNil

`func (o *CreateCluster) SetApiServerIngressRulesNil(b bool)`

 SetApiServerIngressRulesNil sets the value for ApiServerIngressRules to be an explicit nil

### UnsetApiServerIngressRules
`func (o *CreateCluster) UnsetApiServerIngressRules()`

UnsetApiServerIngressRules ensures that no value is present for ApiServerIngressRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


