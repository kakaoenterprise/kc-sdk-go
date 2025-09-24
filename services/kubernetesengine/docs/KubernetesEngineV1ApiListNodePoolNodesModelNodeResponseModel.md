# KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCordon** | **bool** | 노드가 cordon(스케줄링 비활성) 상태인지 여부 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Flavor** | **string** | 인스턴스 유형 | 
**Id** | **string** | 노드의 고유 ID | 
**Image** | [**ImageInfoResponseModel**](ImageInfoResponseModel.md) | 노드에서 사용 중인 이미지 | 
**Ip** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | 노드 이름 | 
**NodePoolName** | **string** | 노드 풀 이름 | 
**SshKeyName** | **string** | SSH 키 이름 | 
**Status** | [**KubernetesEngineV1ApiListNodePoolNodesModelStatusInfoResponseModel**](KubernetesEngineV1ApiListNodePoolNodesModelStatusInfoResponseModel.md) | 노드 상태 | 
**FailureMessage** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Version** | **string** | 노드의 Kubernetes 버전 | 
**VolumeSize** | **int32** | 노드의 루트 볼륨 크기(단위: GiB) | 
**VpcInfo** | [**KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel**](KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel.md) | 노드가 속한 VPC 정보 | 
**IsHyperThreading** | **bool** | 하이퍼스레딩 사용 여부 | 

## Methods

### NewKubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel

`func NewKubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel(isCordon bool, createdAt time.Time, flavor string, id string, image ImageInfoResponseModel, name string, nodePoolName string, sshKeyName string, status KubernetesEngineV1ApiListNodePoolNodesModelStatusInfoResponseModel, updatedAt time.Time, version string, volumeSize int32, vpcInfo KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel, isHyperThreading bool, ) *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel`

NewKubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel instantiates a new KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModelWithDefaults() *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel`

NewKubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCordon

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetIsCordon() bool`

GetIsCordon returns the IsCordon field if non-nil, zero value otherwise.

### GetIsCordonOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetIsCordonOk() (*bool, bool)`

GetIsCordonOk returns a tuple with the IsCordon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCordon

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetIsCordon(v bool)`

SetIsCordon sets IsCordon field to given value.


### GetCreatedAt

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFlavor

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetId

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetImage() ImageInfoResponseModel`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetImageOk() (*ImageInfoResponseModel, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetImage(v ImageInfoResponseModel)`

SetImage sets Image field to given value.


### GetIp

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetName

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetNodePoolName

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetNodePoolName() string`

GetNodePoolName returns the NodePoolName field if non-nil, zero value otherwise.

### GetNodePoolNameOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetNodePoolNameOk() (*string, bool)`

GetNodePoolNameOk returns a tuple with the NodePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePoolName

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetNodePoolName(v string)`

SetNodePoolName sets NodePoolName field to given value.


### GetSshKeyName

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetStatus

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetStatus() KubernetesEngineV1ApiListNodePoolNodesModelStatusInfoResponseModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetStatusOk() (*KubernetesEngineV1ApiListNodePoolNodesModelStatusInfoResponseModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetStatus(v KubernetesEngineV1ApiListNodePoolNodesModelStatusInfoResponseModel)`

SetStatus sets Status field to given value.


### GetFailureMessage

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### SetFailureMessageNil

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetUpdatedAt

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVolumeSize

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetVpcInfo

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetVpcInfo() KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetVpcInfoOk() (*KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetVpcInfo(v KubernetesEngineV1ApiListNodePoolNodesModelVpcInfoResponseModel)`

SetVpcInfo sets VpcInfo field to given value.


### GetIsHyperThreading

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *KubernetesEngineV1ApiListNodePoolNodesModelNodeResponseModel) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


