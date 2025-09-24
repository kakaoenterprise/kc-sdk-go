# KubernetesEngineV1ApiGetClusterModelClusterResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAllocateFip** | **bool** | 퍼블릭 IP 할당 여부 | 
**ApiVersion** | **string** | API 버전 | 
**Network** | Pointer to [**NullableClusterNetworkResponseModel**](ClusterNetworkResponseModel.md) |  | [optional] 
**ControlPlaneEndpoint** | [**ControlPlaneEndpointResponseModel**](ControlPlaneEndpointResponseModel.md) | 컨트롤 플레인 접속 정보 | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatorInfo** | [**CreatorInfoResponseModel**](CreatorInfoResponseModel.md) | 생성자 정보 | 
**Version** | [**OMTInfoResponseModel**](OMTInfoResponseModel.md) | 클러스터 Kubernetes 버전 정보 | 
**Description** | **string** | 클러스터에 대한 설명 | 
**Id** | **string** | 클러스터의 고유 ID | 
**Name** | **string** | 클러스터 이름 | 
**Status** | [**KubernetesEngineV1ApiGetClusterModelStatusResponseModel**](KubernetesEngineV1ApiGetClusterModelStatusResponseModel.md) | 클러스터 상태 | 
**FailureMessage** | Pointer to **NullableString** |  | [optional] 
**IsUpgradable** | **bool** | 클러스터 업그레이드 가능 여부 | 
**VpcInfo** | [**KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel**](KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel.md) | 클러스터가 속한 VPC 정보 | 

## Methods

### NewKubernetesEngineV1ApiGetClusterModelClusterResponseModel

`func NewKubernetesEngineV1ApiGetClusterModelClusterResponseModel(isAllocateFip bool, apiVersion string, controlPlaneEndpoint ControlPlaneEndpointResponseModel, creatorInfo CreatorInfoResponseModel, version OMTInfoResponseModel, description string, id string, name string, status KubernetesEngineV1ApiGetClusterModelStatusResponseModel, isUpgradable bool, vpcInfo KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel, ) *KubernetesEngineV1ApiGetClusterModelClusterResponseModel`

NewKubernetesEngineV1ApiGetClusterModelClusterResponseModel instantiates a new KubernetesEngineV1ApiGetClusterModelClusterResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiGetClusterModelClusterResponseModelWithDefaults

`func NewKubernetesEngineV1ApiGetClusterModelClusterResponseModelWithDefaults() *KubernetesEngineV1ApiGetClusterModelClusterResponseModel`

NewKubernetesEngineV1ApiGetClusterModelClusterResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiGetClusterModelClusterResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAllocateFip

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetIsAllocateFip() bool`

GetIsAllocateFip returns the IsAllocateFip field if non-nil, zero value otherwise.

### GetIsAllocateFipOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetIsAllocateFipOk() (*bool, bool)`

GetIsAllocateFipOk returns a tuple with the IsAllocateFip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllocateFip

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetIsAllocateFip(v bool)`

SetIsAllocateFip sets IsAllocateFip field to given value.


### GetApiVersion

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetNetwork

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetNetwork() ClusterNetworkResponseModel`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetNetworkOk() (*ClusterNetworkResponseModel, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetNetwork(v ClusterNetworkResponseModel)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetControlPlaneEndpoint

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetControlPlaneEndpoint() ControlPlaneEndpointResponseModel`

GetControlPlaneEndpoint returns the ControlPlaneEndpoint field if non-nil, zero value otherwise.

### GetControlPlaneEndpointOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetControlPlaneEndpointOk() (*ControlPlaneEndpointResponseModel, bool)`

GetControlPlaneEndpointOk returns a tuple with the ControlPlaneEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneEndpoint

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetControlPlaneEndpoint(v ControlPlaneEndpointResponseModel)`

SetControlPlaneEndpoint sets ControlPlaneEndpoint field to given value.


### GetCreatedAt

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorInfo

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetCreatorInfo() CreatorInfoResponseModel`

GetCreatorInfo returns the CreatorInfo field if non-nil, zero value otherwise.

### GetCreatorInfoOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetCreatorInfoOk() (*CreatorInfoResponseModel, bool)`

GetCreatorInfoOk returns a tuple with the CreatorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorInfo

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetCreatorInfo(v CreatorInfoResponseModel)`

SetCreatorInfo sets CreatorInfo field to given value.


### GetVersion

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetVersion() OMTInfoResponseModel`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetVersionOk() (*OMTInfoResponseModel, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetVersion(v OMTInfoResponseModel)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetStatus() KubernetesEngineV1ApiGetClusterModelStatusResponseModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetStatusOk() (*KubernetesEngineV1ApiGetClusterModelStatusResponseModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetStatus(v KubernetesEngineV1ApiGetClusterModelStatusResponseModel)`

SetStatus sets Status field to given value.


### GetFailureMessage

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### SetFailureMessageNil

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetIsUpgradable

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetIsUpgradable() bool`

GetIsUpgradable returns the IsUpgradable field if non-nil, zero value otherwise.

### GetIsUpgradableOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetIsUpgradableOk() (*bool, bool)`

GetIsUpgradableOk returns a tuple with the IsUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradable

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetIsUpgradable(v bool)`

SetIsUpgradable sets IsUpgradable field to given value.


### GetVpcInfo

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetVpcInfo() KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) GetVpcInfoOk() (*KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *KubernetesEngineV1ApiGetClusterModelClusterResponseModel) SetVpcInfo(v KubernetesEngineV1ApiGetClusterModelVPCInfoResponseModel)`

SetVpcInfo sets VpcInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


