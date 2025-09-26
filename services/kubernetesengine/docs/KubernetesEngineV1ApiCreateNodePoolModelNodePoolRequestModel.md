# KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 대상 노드 풀 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**FlavorId** | **string** | 인스턴스 유형 ID | 
**VolumeSize** | Pointer to **NullableInt32** |  | [optional] 
**NodeCount** | **int32** | 노드 풀의 노드 수 | 
**SshKeyName** | **string** | SSH 키 이름 | 
**Labels** | Pointer to [**[]LabelRequestModel**](LabelRequestModel.md) |  | [optional] 
**Taints** | Pointer to [**[]TaintRequestModel**](TaintRequestModel.md) |  | [optional] 
**UserData** | Pointer to **NullableString** |  | [optional] 
**VpcInfo** | [**VpcInfoRequestModel**](VpcInfoRequestModel.md) | 노드 풀을 생성할 VPC 정보 | 
**ImageId** | **string** | 이미지의 고유 ID | 
**IsHyperThreading** | Pointer to **NullableBool** |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel

`func NewKubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel(name string, flavorId string, nodeCount int32, sshKeyName string, vpcInfo VpcInfoRequestModel, imageId string, ) *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel`

NewKubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel instantiates a new KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModelWithDefaults

`func NewKubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModelWithDefaults() *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel`

NewKubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModelWithDefaults instantiates a new KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFlavorId

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetVolumeSize

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### SetVolumeSizeNil

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetVolumeSizeNil(b bool)`

 SetVolumeSizeNil sets the value for VolumeSize to be an explicit nil

### UnsetVolumeSize
`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) UnsetVolumeSize()`

UnsetVolumeSize ensures that no value is present for VolumeSize, not even an explicit nil
### GetNodeCount

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetSshKeyName

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetSshKeyName() string`

GetSshKeyName returns the SshKeyName field if non-nil, zero value otherwise.

### GetSshKeyNameOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetSshKeyNameOk() (*string, bool)`

GetSshKeyNameOk returns a tuple with the SshKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyName

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetSshKeyName(v string)`

SetSshKeyName sets SshKeyName field to given value.


### GetLabels

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetLabels() []LabelRequestModel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetLabelsOk() (*[]LabelRequestModel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetLabels(v []LabelRequestModel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTaints

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetTaints() []TaintRequestModel`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetTaintsOk() (*[]TaintRequestModel, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetTaints(v []TaintRequestModel)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### SetTaintsNil

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetTaintsNil(b bool)`

 SetTaintsNil sets the value for Taints to be an explicit nil

### UnsetTaints
`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) UnsetTaints()`

UnsetTaints ensures that no value is present for Taints, not even an explicit nil
### GetUserData

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVpcInfo

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetVpcInfo() VpcInfoRequestModel`

GetVpcInfo returns the VpcInfo field if non-nil, zero value otherwise.

### GetVpcInfoOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetVpcInfoOk() (*VpcInfoRequestModel, bool)`

GetVpcInfoOk returns a tuple with the VpcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcInfo

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetVpcInfo(v VpcInfoRequestModel)`

SetVpcInfo sets VpcInfo field to given value.


### GetImageId

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetIsHyperThreading

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.

### HasIsHyperThreading

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) HasIsHyperThreading() bool`

HasIsHyperThreading returns a boolean if a field has been set.

### SetIsHyperThreadingNil

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetIsHyperThreadingNil(b bool)`

 SetIsHyperThreadingNil sets the value for IsHyperThreading to be an explicit nil

### UnsetIsHyperThreading
`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) UnsetIsHyperThreading()`

UnsetIsHyperThreading ensures that no value is present for IsHyperThreading, not even an explicit nil
### GetSecurityGroups

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *KubernetesEngineV1ApiCreateNodePoolModelNodePoolRequestModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


