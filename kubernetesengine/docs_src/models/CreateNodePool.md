# CreateNodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 대상 노드 풀 이름 | 
**Description** | Pointer to **NullableString** | 노드 풀에 대한 설명 | [optional] 
**FlavorId** | **string** | 인스턴스 유형 ID - [List instance types](/openapi/bcs/list-instance-types)에서 확인 | 
**VolumeSize** | Pointer to **NullableInt32** | 노드 풀의 루트 볼륨 크기(단위: GiB) - vm, gpu 노드 풀 생성 시 정의 필요 - vm, gpu 노드 풀 생성 시, &#x60;volume_size&#x60;를 입력하지 않으면 &#x60;50&#x60; (기본값)으로 생성됨 | [optional] 
**NodeCount** | **int32** | 노드 풀의 노드 수 | 
**SshKeyName** | **string** | SSH 키 이름 | 
**Labels** | Pointer to [**[]LabelRequest**](LabelRequest.md) | 노드 풀에 적용할 레이블 목록 (Key-Value 쌍) | [optional] 
**Taints** | Pointer to [**[]TaintRequest**](TaintRequest.md) | 노드 풀에 적용할 테인트 목록 | [optional] 
**UserData** | Pointer to **NullableString** | 노드 풀 내 노드 생성 시 실행할 사용자 스크립트(base64 인코딩 필요) | [optional] 
**VpcInfo** | [**VpcInfoRequest**](VpcInfoRequest.md) | 노드 풀을 생성할 VPC 정보 | 
**ImageId** | **string** | 이미지의 고유 ID - [List node pool images](/openapi/container-pack/k8se/list-node-pool-images)에서 확인 | 
**IsHyperThreading** | Pointer to **NullableBool** | 하이퍼스레딩 사용 여부 - &#x60;true&#x60;: 노드에서 하이퍼스레딩을 활성화하여 물리 코어당 2개의 vCPU를 인식 - &#x60;false&#x60;: 하이퍼스레딩을 비활성화하여 물리 코어 수와 동일한 vCPU만 인식 | [optional] 
**SecurityGroups** | Pointer to **[]string** | 연결할 보안 그룹 ID 목록 - [List security groups](/openapi/networking/vpc/list-security-groups)에서 확인 | [optional] 

## Methods

### NewCreateNodePool

`func NewCreateNodePool(name string, flavorId string, nodeCount int32, sshKeyName string, vpcInfo VpcInfoRequest, imageId string, ) *CreateNodePool`

NewCreateNodePool instantiates a new CreateNodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNodePoolWithDefaults

`func NewCreateNodePoolWithDefaults() *CreateNodePool`

NewCreateNodePoolWithDefaults instantiates a new CreateNodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNodePool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateNodePool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNodePool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNodePool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNodePool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateNodePool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateNodePool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFlavorId

`func (o *CreateNodePool) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *CreateNodePool) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *CreateNodePool) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetVolumeSize

`func (o *CreateNodePool) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *CreateNodePool) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *CreateNodePool) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *CreateNodePool) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### SetVolumeSizeNil

`func (o *CreateNodePool) SetVolumeSizeNil(b bool)`

 SetVolumeSizeNil sets the value for VolumeSize to be an explicit nil

### UnsetVolumeSize
`func (o *CreateNodePool) UnsetVolumeSize()`

UnsetVolumeSize ensures that no value is present for VolumeSize, not even an explicit nil
### GetNodeCount

`func (o *CreateNodePool) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CreateNodePool) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CreateNodePool) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetSshKeyName

`func (o *CreateNodePool) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *CreateNodePool) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *CreateNodePool) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetLabels

`func (o *CreateNodePool) GetLabels() []LabelRequest`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateNodePool) GetLabelsOk() (*[]LabelRequest, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateNodePool) SetLabels(v []LabelRequest)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateNodePool) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *CreateNodePool) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *CreateNodePool) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTaints

`func (o *CreateNodePool) GetTaints() []TaintRequest`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *CreateNodePool) GetTaintsOk() (*[]TaintRequest, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *CreateNodePool) SetTaints(v []TaintRequest)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *CreateNodePool) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### SetTaintsNil

`func (o *CreateNodePool) SetTaintsNil(b bool)`

 SetTaintsNil sets the value for Taints to be an explicit nil

### UnsetTaints
`func (o *CreateNodePool) UnsetTaints()`

UnsetTaints ensures that no value is present for Taints, not even an explicit nil
### GetUserData

`func (o *CreateNodePool) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateNodePool) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateNodePool) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateNodePool) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *CreateNodePool) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *CreateNodePool) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVpcInfo

`func (o *CreateNodePool) GetVpcInfo() VpcInfoRequest`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *CreateNodePool) GetVpcInfoOk() (*VpcInfoRequest, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *CreateNodePool) SetVpcInfo(v VpcInfoRequest)`

SetVpcInfo sets VpcInfo field to given value.


### GetImageId

`func (o *CreateNodePool) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateNodePool) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateNodePool) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetIsHyperThreading

`func (o *CreateNodePool) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *CreateNodePool) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *CreateNodePool) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.

### HasIsHyperThreading

`func (o *CreateNodePool) HasIsHyperThreading() bool`

HasIsHyperThreading returns a boolean if a field has been set.

### SetIsHyperThreadingNil

`func (o *CreateNodePool) SetIsHyperThreadingNil(b bool)`

 SetIsHyperThreadingNil sets the value for IsHyperThreading to be an explicit nil

### UnsetIsHyperThreading
`func (o *CreateNodePool) UnsetIsHyperThreading()`

UnsetIsHyperThreading ensures that no value is present for IsHyperThreading, not even an explicit nil
### GetSecurityGroups

`func (o *CreateNodePool) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateNodePool) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateNodePool) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateNodePool) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CreateNodePool) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CreateNodePool) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


