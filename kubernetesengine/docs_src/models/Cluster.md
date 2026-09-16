# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAllocateFip** | **bool** | 퍼블릭 IP 할당 여부 | 
**ApiVersion** | **string** | API 버전 | 
**Network** | Pointer to [**NullableClusterNetwork**](ClusterNetwork.md) | 클러스터 네트워크 구성 정보 | [optional] 
**ApiServerIngressRules** | [**[]ApiServerIngressRule**](ApiServerIngressRule.md) | API 서버 접근을 허용하는 CIDR 목록 | 
**ControlPlaneEndpoint** | [**ControlPlaneEndpoint**](ControlPlaneEndpoint.md) | 컨트롤 플레인 접속 정보 | 
**CreatedAt** | Pointer to **NullableTime** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**CreatorInfo** | [**CreatorInfo**](CreatorInfo.md) | 생성자 정보 | 
**Version** | [**OMTInfo**](OMTInfo.md) | 클러스터 Kubernetes 버전 정보 | 
**Description** | **string** | 클러스터에 대한 설명 | 
**Id** | **string** | 클러스터의 고유 ID | 
**Name** | **string** | 클러스터 이름 | 
**Status** | [**Status**](Status.md) | 클러스터 상태 | 
**FailureMessage** | Pointer to **NullableString** | 클러스터의 상태가 &#x60;Failed&#x60;로 변경된 경우 실패 메시지 | [optional] 
**IsUpgradable** | **bool** | 업그레이드 가능 여부 | 
**VpcInfo** | [**VpcInfo**](VpcInfo.md) | 클러스터가 속한 VPC 정보 | 

## Methods

### NewCluster

`func NewCluster(isAllocateFip bool, apiVersion string, apiServerIngressRules []ApiServerIngressRule, controlPlaneEndpoint ControlPlaneEndpoint, creatorInfo CreatorInfo, version OMTInfo, description string, id string, name string, status Status, isUpgradable bool, vpcInfo VpcInfo, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAllocateFip

`func (o *Cluster) GetIsAllocateFip() bool`

GetIsAllocateFip returns the IsAllocateFip field if non-nil, zero value otherwise.

### GetIsAllocateFipOk

`func (o *Cluster) GetIsAllocateFipOk() (*bool, bool)`

GetIsAllocateFipOk returns a tuple with the IsAllocateFip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllocateFip

`func (o *Cluster) SetIsAllocateFip(v bool)`

SetIsAllocateFip sets IsAllocateFip field to given value.


### GetApiVersion

`func (o *Cluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Cluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Cluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetNetwork

`func (o *Cluster) GetNetwork() ClusterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Cluster) GetNetworkOk() (*ClusterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Cluster) SetNetwork(v ClusterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Cluster) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *Cluster) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *Cluster) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetApiServerIngressRules

`func (o *Cluster) GetApiServerIngressRules() []ApiServerIngressRule`

GetApiServerIngressRules returns the ApiServerIngressRules field if non-nil, zero value otherwise.

### GetApiServerIngressRulesOk

`func (o *Cluster) GetApiServerIngressRulesOk() (*[]ApiServerIngressRule, bool)`

GetApiServerIngressRulesOk returns a tuple with the ApiServerIngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerIngressRules

`func (o *Cluster) SetApiServerIngressRules(v []ApiServerIngressRule)`

SetApiServerIngressRules sets ApiServerIngressRules field to given value.


### GetControlPlaneEndpoint

`func (o *Cluster) GetControlPlaneEndpoint() ControlPlaneEndpoint`

GetControlPlaneEndpoint returns the ControlPlaneEndpoint field if non-nil, zero value otherwise.

### GetControlPlaneEndpointOk

`func (o *Cluster) GetControlPlaneEndpointOk() (*ControlPlaneEndpoint, bool)`

GetControlPlaneEndpointOk returns a tuple with the ControlPlaneEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneEndpoint

`func (o *Cluster) SetControlPlaneEndpoint(v ControlPlaneEndpoint)`

SetControlPlaneEndpoint sets ControlPlaneEndpoint field to given value.


### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Cluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Cluster) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Cluster) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorInfo

`func (o *Cluster) GetCreatorInfo() CreatorInfo`

GetCreatorInfo returns the CreatorInfo field if non-nil, zero value otherwise.

### GetCreatorInfoOk

`func (o *Cluster) GetCreatorInfoOk() (*CreatorInfo, bool)`

GetCreatorInfoOk returns a tuple with the CreatorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorInfo

`func (o *Cluster) SetCreatorInfo(v CreatorInfo)`

SetCreatorInfo sets CreatorInfo field to given value.


### GetVersion

`func (o *Cluster) GetVersion() OMTInfo`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*OMTInfo, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v OMTInfo)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Cluster) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetFailureMessage

`func (o *Cluster) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *Cluster) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *Cluster) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *Cluster) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### SetFailureMessageNil

`func (o *Cluster) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *Cluster) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetIsUpgradable

`func (o *Cluster) GetIsUpgradable() bool`

GetIsUpgradable returns the IsUpgradable field if non-nil, zero value otherwise.

### GetIsUpgradableOk

`func (o *Cluster) GetIsUpgradableOk() (*bool, bool)`

GetIsUpgradableOk returns a tuple with the IsUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradable

`func (o *Cluster) SetIsUpgradable(v bool)`

SetIsUpgradable sets IsUpgradable field to given value.


### GetVpcInfo

`func (o *Cluster) GetVpcInfo() VpcInfo`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *Cluster) GetVpcInfoOk() (*VpcInfo, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *Cluster) SetVpcInfo(v VpcInfo)`

SetVpcInfo sets VpcInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


