# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCordon** | **bool** | 노드가 cordon(스케줄링 비활성) 상태인지 여부 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
**Flavor** | **string** | 인스턴스 유형 | 
**Id** | **string** | 노드의 고유 ID | 
**Image** | [**ImageInfo**](ImageInfo.md) | 노드에서 사용 중인 이미지 | 
**Ip** | Pointer to **NullableString** | 노드의 IP 주소 | [optional] 
**Name** | **string** | 노드 이름 | 
**NodePoolName** | **string** | 노드 풀 이름 | 
**SshKeyName** | **string** | SSH 키 이름 | 
**Status** | [**StatusInfo**](StatusInfo.md) | 노드 상태 | 
**FailureMessage** | Pointer to **NullableString** | 노드의 상태가 &#x60;Failed&#x60;로 변경된 경우 실패 메시지 | [optional] 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
**Version** | **string** | 노드의 Kubernetes 버전 | 
**VolumeSize** | **int32** | 노드의 루트 볼륨 크기(단위: GiB) | 
**VpcInfo** | [**VpcInfo**](VpcInfo.md) | 노드가 속한 VPC 정보 | 
**IsHyperThreading** | **bool** | 하이퍼스레딩 사용 여부 | 

## Methods

### NewNode

`func NewNode(isCordon bool, createdAt time.Time, flavor string, id string, image ImageInfo, name string, nodePoolName string, sshKeyName string, status StatusInfo, updatedAt time.Time, version string, volumeSize int32, vpcInfo VpcInfo, isHyperThreading bool, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCordon

`func (o *Node) GetIsCordon() bool`

GetIsCordon returns the IsCordon field if non-nil, zero value otherwise.

### GetIsCordonOk

`func (o *Node) GetIsCordonOk() (*bool, bool)`

GetIsCordonOk returns a tuple with the IsCordon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCordon

`func (o *Node) SetIsCordon(v bool)`

SetIsCordon sets IsCordon field to given value.


### GetCreatedAt

`func (o *Node) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Node) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Node) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFlavor

`func (o *Node) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Node) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Node) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### GetId

`func (o *Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *Node) GetImage() ImageInfo`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Node) GetImageOk() (*ImageInfo, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Node) SetImage(v ImageInfo)`

SetImage sets Image field to given value.


### GetIp

`func (o *Node) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Node) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Node) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Node) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *Node) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *Node) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.


### GetNodePoolName

`func (o *Node) GetNodePoolName() string`

GetNodePoolName returns the NodePoolName field if non-nil, zero value otherwise.

### GetNodePoolNameOk

`func (o *Node) GetNodePoolNameOk() (*string, bool)`

GetNodePoolNameOk returns a tuple with the NodePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePoolName

`func (o *Node) SetNodePoolName(v string)`

SetNodePoolName sets NodePoolName field to given value.


### GetSshKeyName

`func (o *Node) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *Node) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *Node) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetStatus

`func (o *Node) GetStatus() StatusInfo`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Node) GetStatusOk() (*StatusInfo, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Node) SetStatus(v StatusInfo)`

SetStatus sets Status field to given value.


### GetFailureMessage

`func (o *Node) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *Node) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *Node) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *Node) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### SetFailureMessageNil

`func (o *Node) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *Node) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetUpdatedAt

`func (o *Node) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Node) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Node) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *Node) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Node) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Node) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVolumeSize

`func (o *Node) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *Node) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *Node) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetVpcInfo

`func (o *Node) GetVpcInfo() VpcInfo`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *Node) GetVpcInfoOk() (*VpcInfo, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *Node) SetVpcInfo(v VpcInfo)`

SetVpcInfo sets VpcInfo field to given value.


### GetIsHyperThreading

`func (o *Node) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *Node) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *Node) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


