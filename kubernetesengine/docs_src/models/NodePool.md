# NodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBareMetal** | **bool** | 베어메탈 유형 노드 풀 여부 | 
**ClusterName** | **string** | 대상 클러스터 이름 | 
**IsCordon** | **bool** | 노드 풀 내 전체 노드가 cordon(스케줄링 비활성) 상태인지 여부 | 
**NodeCount** | **int32** | 노드 풀의 노드 수 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
**Description** | **string** | 노드 풀에 대한 설명 | 
**FailureMessage** | Pointer to **NullableString** | 노드 풀 내 노드의 상태가 &#x60;Failed&#x60;로 변경된 경우 실패 메시지 | [optional] 
**Flavor** | **string** | 인스턴스 유형 이름 | 
**FlavorId** | **string** | 인스턴스 유형 ID | 
**IsGpu** | **bool** | GPU 유형 노드 풀 여부 | 
**Id** | **string** | 노드 풀의 고유 ID | 
**Image** | [**ImageInfo**](ImageInfo.md) | 노드 이미지 정보 | 
**Labels** | [**[]LabelInfo**](LabelInfo.md) | 노드 풀에 적용된 레이블 목록 | 
**Name** | **string** | 노드 풀 이름 | 
**SecurityGroups** | **[]string** | 연결된 보안 그룹 목록 | 
**SshKeyName** | **string** | SSH 키 이름 | 
**Status** | [**NodePoolStatusInfo**](NodePoolStatusInfo.md) | 노드 풀 상태 | 
**Taints** | [**[]TaintInfo**](TaintInfo.md) | 노드 풀이 가진 테인트 목록 | 
**IsUpgradable** | **bool** | 업그레이드 가능 여부 | 
**UserData** | Pointer to **NullableString** | 노드 풀 설정된 사용자 스크립트 (base64 인코딩된 값) | [optional] 
**Version** | **string** | 노드 풀의 Kubernetes 버전 | 
**VolumeSize** | **int32** | 노드의 루트 볼륨 크기(단위: GiB) | 
**VpcInfo** | [**VpcInfo**](VpcInfo.md) | 노드 풀이 속한 VPC 정보 | 
**Autoscaling** | [**Autoscaling**](Autoscaling.md) | 노드 풀의 리소스 기반 오토스케일링 설정 정보 | 
**IsHyperThreading** | **bool** | 하이퍼스레딩 사용 여부 | 

## Methods

### NewNodePool

`func NewNodePool(isBareMetal bool, clusterName string, isCordon bool, nodeCount int32, createdAt time.Time, description string, flavor string, flavorId string, isGpu bool, id string, image ImageInfo, labels []LabelInfo, name string, securityGroups []string, sshKeyName string, status NodePoolStatusInfo, taints []TaintInfo, isUpgradable bool, version string, volumeSize int32, vpcInfo VpcInfo, autoscaling Autoscaling, isHyperThreading bool, ) *NodePool`

NewNodePool instantiates a new NodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolWithDefaults

`func NewNodePoolWithDefaults() *NodePool`

NewNodePoolWithDefaults instantiates a new NodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBareMetal

`func (o *NodePool) GetIsBareMetal() bool`

GetIsBareMetal returns the IsBareMetal field if non-nil, zero value otherwise.

### GetIsBareMetalOk

`func (o *NodePool) GetIsBareMetalOk() (*bool, bool)`

GetIsBareMetalOk returns a tuple with the IsBareMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBareMetal

`func (o *NodePool) SetIsBareMetal(v bool)`

SetIsBareMetal sets IsBareMetal field to given value.


### GetClusterName

`func (o *NodePool) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *NodePool) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *NodePool) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetIsCordon

`func (o *NodePool) GetIsCordon() bool`

GetIsCordon returns the IsCordon field if non-nil, zero value otherwise.

### GetIsCordonOk

`func (o *NodePool) GetIsCordonOk() (*bool, bool)`

GetIsCordonOk returns a tuple with the IsCordon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCordon

`func (o *NodePool) SetIsCordon(v bool)`

SetIsCordon sets IsCordon field to given value.


### GetNodeCount

`func (o *NodePool) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *NodePool) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *NodePool) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetCreatedAt

`func (o *NodePool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodePool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodePool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *NodePool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodePool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodePool) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFailureMessage

`func (o *NodePool) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *NodePool) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *NodePool) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *NodePool) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### SetFailureMessageNil

`func (o *NodePool) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *NodePool) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetFlavor

`func (o *NodePool) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *NodePool) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *NodePool) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetFlavorId

`func (o *NodePool) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *NodePool) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *NodePool) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetIsGpu

`func (o *NodePool) GetIsGpu() bool`

GetIsGpu returns the IsGpu field if non-nil, zero value otherwise.

### GetIsGpuOk

`func (o *NodePool) GetIsGpuOk() (*bool, bool)`

GetIsGpuOk returns a tuple with the IsGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGpu

`func (o *NodePool) SetIsGpu(v bool)`

SetIsGpu sets IsGpu field to given value.


### GetId

`func (o *NodePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePool) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *NodePool) GetImage() ImageInfo`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *NodePool) GetImageOk() (*ImageInfo, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *NodePool) SetImage(v ImageInfo)`

SetImage sets Image field to given value.


### GetLabels

`func (o *NodePool) GetLabels() []LabelInfo`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NodePool) GetLabelsOk() (*[]LabelInfo, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NodePool) SetLabels(v []LabelInfo)`

SetLabels sets Labels field to given value.


### GetName

`func (o *NodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePool) SetName(v string)`

SetName sets Name field to given value.


### GetSecurityGroups

`func (o *NodePool) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *NodePool) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *NodePool) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetSshKeyName

`func (o *NodePool) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *NodePool) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *NodePool) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetStatus

`func (o *NodePool) GetStatus() NodePoolStatusInfo`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodePool) GetStatusOk() (*NodePoolStatusInfo, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodePool) SetStatus(v NodePoolStatusInfo)`

SetStatus sets Status field to given value.


### GetTaints

`func (o *NodePool) GetTaints() []TaintInfo`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *NodePool) GetTaintsOk() (*[]TaintInfo, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *NodePool) SetTaints(v []TaintInfo)`

SetTaints sets Taints field to given value.


### GetIsUpgradable

`func (o *NodePool) GetIsUpgradable() bool`

GetIsUpgradable returns the IsUpgradable field if non-nil, zero value otherwise.

### GetIsUpgradableOk

`func (o *NodePool) GetIsUpgradableOk() (*bool, bool)`

GetIsUpgradableOk returns a tuple with the IsUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradable

`func (o *NodePool) SetIsUpgradable(v bool)`

SetIsUpgradable sets IsUpgradable field to given value.


### GetUserData

`func (o *NodePool) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *NodePool) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *NodePool) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *NodePool) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *NodePool) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *NodePool) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVersion

`func (o *NodePool) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodePool) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodePool) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVolumeSize

`func (o *NodePool) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *NodePool) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *NodePool) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetVpcInfo

`func (o *NodePool) GetVpcInfo() VpcInfo`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *NodePool) GetVpcInfoOk() (*VpcInfo, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *NodePool) SetVpcInfo(v VpcInfo)`

SetVpcInfo sets VpcInfo field to given value.


### GetAutoscaling

`func (o *NodePool) GetAutoscaling() Autoscaling`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *NodePool) GetAutoscalingOk() (*Autoscaling, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *NodePool) SetAutoscaling(v Autoscaling)`

SetAutoscaling sets Autoscaling field to given value.


### GetIsHyperThreading

`func (o *NodePool) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *NodePool) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *NodePool) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


