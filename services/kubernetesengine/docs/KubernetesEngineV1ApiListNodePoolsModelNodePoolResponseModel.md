# KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoscaling** | [**KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel**](KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel.md) | 노드 풀의 리소스 기반 오토스케일링 설정 정보 | 
**IsBareMetal** | **bool** | 베어메탈 유형 노드 풀 여부 | 
**ClusterName** | **string** | 대상 클러스터 이름 | 
**IsCordon** | **bool** | 노드 풀 내 전체 노드가 cordon(스케줄링 비활성) 상태인지 여부 | 
**NodeCount** | **int32** | 노드 풀의 노드 수 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Description** | **string** | 노드 풀에 대한 설명 | 
**FailureMessage** | Pointer to **NullableString** |  | [optional] 
**Flavor** | **string** | 인스턴스 유형 이름 | 
**FlavorId** | **string** | 인스턴스 유형 ID | 
**IsGpu** | **bool** | GPU 유형 노드 풀 여부 | 
**Id** | **string** | 노드 풀의 고유 ID | 
**Image** | [**ImageInfoResponseModel**](ImageInfoResponseModel.md) | 노드 이미지 정보 | 
**Labels** | [**[]LabelInfoResponseModel**](LabelInfoResponseModel.md) | 노드 풀에 적용된 레이블 목록 | 
**Name** | **string** | 노드 풀 이름 | 
**SecurityGroups** | **[]string** | 연결된 보안 그룹 목록 | 
**SshKeyName** | **string** | SSH 키 이름 | 
**Status** | [**KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel**](KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel.md) | 노드 풀 상태 | 
**Taints** | [**[]KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel**](KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel.md) | 노드 풀이 가진 테인트 목록 | 
**IsUpgradable** | **bool** | 업그레이드 가능 여부 | 
**UserData** | Pointer to **NullableString** |  | [optional] 
**Version** | **string** | 노드 풀의 Kubernetes 버전 | 
**VolumeSize** | **int32** | 노드의 루트 볼륨 크기(단위: GiB) | 
**VpcInfo** | [**KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel**](KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel.md) | 노드 풀이 속한 VPC 정보 | 
**IsHyperThreading** | **bool** | 하이퍼스레딩 사용 여부 | 

## Methods

### NewKubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel

`func NewKubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel(autoscaling KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel, isBareMetal bool, clusterName string, isCordon bool, nodeCount int32, createdAt time.Time, description string, flavor string, flavorId string, isGpu bool, id string, image ImageInfoResponseModel, labels []LabelInfoResponseModel, name string, securityGroups []string, sshKeyName string, status KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel, taints []KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel, isUpgradable bool, version string, volumeSize int32, vpcInfo KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel, isHyperThreading bool, ) *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel instantiates a new KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModelWithDefaults() *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscaling

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetAutoscaling() KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetAutoscalingOk() (*KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetAutoscaling(v KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel)`

SetAutoscaling sets Autoscaling field to given value.


### GetIsBareMetal

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsBareMetal() bool`

GetIsBareMetal returns the IsBareMetal field if non-nil, zero value otherwise.

### GetIsBareMetalOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsBareMetalOk() (*bool, bool)`

GetIsBareMetalOk returns a tuple with the IsBareMetal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBareMetal

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetIsBareMetal(v bool)`

SetIsBareMetal sets IsBareMetal field to given value.


### GetClusterName

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetIsCordon

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsCordon() bool`

GetIsCordon returns the IsCordon field if non-nil, zero value otherwise.

### GetIsCordonOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsCordonOk() (*bool, bool)`

GetIsCordonOk returns a tuple with the IsCordon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCordon

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetIsCordon(v bool)`

SetIsCordon sets IsCordon field to given value.


### GetNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetCreatedAt

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFailureMessage

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### SetFailureMessageNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetFlavor

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetFlavorId

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetIsGpu

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsGpu() bool`

GetIsGpu returns the IsGpu field if non-nil, zero value otherwise.

### GetIsGpuOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsGpuOk() (*bool, bool)`

GetIsGpuOk returns a tuple with the IsGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGpu

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetIsGpu(v bool)`

SetIsGpu sets IsGpu field to given value.


### GetId

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetImage() ImageInfoResponseModel`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetImageOk() (*ImageInfoResponseModel, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetImage(v ImageInfoResponseModel)`

SetImage sets Image field to given value.


### GetLabels

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetLabels() []LabelInfoResponseModel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetLabelsOk() (*[]LabelInfoResponseModel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetLabels(v []LabelInfoResponseModel)`

SetLabels sets Labels field to given value.


### GetName

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetSecurityGroups

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetSshKeyName

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetStatus

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetStatus() KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetStatusOk() (*KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetStatus(v KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel)`

SetStatus sets Status field to given value.


### GetTaints

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetTaints() []KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetTaintsOk() (*[]KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetTaints(v []KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel)`

SetTaints sets Taints field to given value.


### GetIsUpgradable

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsUpgradable() bool`

GetIsUpgradable returns the IsUpgradable field if non-nil, zero value otherwise.

### GetIsUpgradableOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsUpgradableOk() (*bool, bool)`

GetIsUpgradableOk returns a tuple with the IsUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradable

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetIsUpgradable(v bool)`

SetIsUpgradable sets IsUpgradable field to given value.


### GetUserData

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVersion

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVolumeSize

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetVpcInfo

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetVpcInfo() KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetVpcInfoOk() (*KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetVpcInfo(v KubernetesEngineV1ApiListNodePoolsModelVpcInfoResponseModel)`

SetVpcInfo sets VpcInfo field to given value.


### GetIsHyperThreading

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *KubernetesEngineV1ApiListNodePoolsModelNodePoolResponseModel) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


