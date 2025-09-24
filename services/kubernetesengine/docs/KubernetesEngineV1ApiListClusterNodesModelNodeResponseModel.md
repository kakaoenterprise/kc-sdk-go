# KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCordon** | **bool** | 노드가 cordon(스케줄링 비활성) 상태인지 여부 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Flavor** | **string** | 인스턴스 유형 | 
**Id** | **string** | 노드의 고유 ID | 
**Image** | [**ImageInfoResponseModel**](ImageInfoResponseModel.md) | 노드 이미지 정보 | 
**Ip** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | 노드 이름 | 
**NodePoolName** | **string** | 노드 풀 이름 | 
**SshKeyName** | **string** | SSH 키 이름 | 
**Status** | [**KubernetesEngineV1ApiListClusterNodesModelStatusInfoResponseModel**](KubernetesEngineV1ApiListClusterNodesModelStatusInfoResponseModel.md) | 노드 상태 | 
**FailureMessage** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Version** | **string** | 노드의 Kubernetes(kubelet) 버전 | 
**VolumeSize** | **int32** | 노드의 루트 볼륨 크기(단위: GiB) | 
**VpcInfo** | [**KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel**](KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel.md) | 노드가 속한 VPC 정보 | 
**IsHyperThreading** | **bool** | 하이퍼스레딩 사용 여부 | 

## Methods

### NewKubernetesEngineV1ApiListClusterNodesModelNodeResponseModel

`func NewKubernetesEngineV1ApiListClusterNodesModelNodeResponseModel(isCordon bool, createdAt time.Time, flavor string, id string, image ImageInfoResponseModel, name string, nodePoolName string, sshKeyName string, status KubernetesEngineV1ApiListClusterNodesModelStatusInfoResponseModel, updatedAt time.Time, version string, volumeSize int32, vpcInfo KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel, isHyperThreading bool, ) *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel`

NewKubernetesEngineV1ApiListClusterNodesModelNodeResponseModel instantiates a new KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListClusterNodesModelNodeResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListClusterNodesModelNodeResponseModelWithDefaults() *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel`

NewKubernetesEngineV1ApiListClusterNodesModelNodeResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCordon

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetIsCordon() bool`

GetIsCordon returns the IsCordon field if non-nil, zero value otherwise.

### GetIsCordonOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetIsCordonOk() (*bool, bool)`

GetIsCordonOk returns a tuple with the IsCordon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCordon

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetIsCordon(v bool)`

SetIsCordon sets IsCordon field to given value.


### GetCreatedAt

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFlavor

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetId

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetImage() ImageInfoResponseModel`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetImageOk() (*ImageInfoResponseModel, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetImage(v ImageInfoResponseModel)`

SetImage sets Image field to given value.


### GetIp

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetName

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetNodePoolName

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetNodePoolName() string`

GetNodePoolName returns the NodePoolName field if non-nil, zero value otherwise.

### GetNodePoolNameOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetNodePoolNameOk() (*string, bool)`

GetNodePoolNameOk returns a tuple with the NodePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePoolName

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetNodePoolName(v string)`

SetNodePoolName sets NodePoolName field to given value.


### GetSshKeyName

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetStatus

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetStatus() KubernetesEngineV1ApiListClusterNodesModelStatusInfoResponseModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetStatusOk() (*KubernetesEngineV1ApiListClusterNodesModelStatusInfoResponseModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetStatus(v KubernetesEngineV1ApiListClusterNodesModelStatusInfoResponseModel)`

SetStatus sets Status field to given value.


### GetFailureMessage

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### SetFailureMessageNil

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetUpdatedAt

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVolumeSize

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetVpcInfo

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetVpcInfo() KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetVpcInfoOk() (*KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetVpcInfo(v KubernetesEngineV1ApiListClusterNodesModelVpcInfoResponseModel)`

SetVpcInfo sets VpcInfo field to given value.


### GetIsHyperThreading

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *KubernetesEngineV1ApiListClusterNodesModelNodeResponseModel) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


